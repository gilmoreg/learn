defmodule KeyVault do
  @moduledoc """
  Documentation for KeyVault module.

  In config.exs:

  config :keyvault,
    azure_client_id: {:system, "<azure_client_id>"},
    azure_client_secret: {:system, System.get_env("AZ_KEYVAULT_SECRET")},
    azure_tenant_id: {:system, System.get_env("AZ_KEYVAULT_TENANT_ID")},
    azure_vault_name: {:system, "<azure_vault_name>"}

  Add KeyVault to your supervisor tree
  """
  use Agent

  def start_link(_opts) do
    Agent.start_link(fn -> Map.new end, name: __MODULE__)
  end

  def get_secret(key) do
    case Agent.get(__MODULE__, &Map.get(&1, key)) do
      nil ->
        {:ok, value} = keyvault_get_secret(key)
        Agent.update(__MODULE__, fn state -> Map.merge(state, %{key => value}) end)
        value
      value -> value
    end
  end

  @spec keyvault_get_secret(String.t) :: String.t
  defp keyvault_get_secret(key) do
    {:ok, auth_header} = get_bearer_token()
    url = "#{keyvault_url()}/secrets/#{key}?api-version=2016-10-01"
    case HTTPoison.get(url, [{"Authorization", auth_header}], http_options()) do
      {:ok, %HTTPoison.Response{status_code: 200, body: body}} ->
        response = Poison.decode!(body)
        {:ok, response["value"]}
      _ ->
        {:error, "Something went wrong"}
    end
  end

  defp keyvault_url do
    {:system, vault_name} = Application.get_env(:keyvault, :azure_vault_name)
    "https://#{vault_name}.vault.azure.net"
  end

  defp get_bearer_token do
    url = auth_url()
    body = auth_body()
    headers = ["Content-Type": "application/x-www-form-urlencoded"]
    case HTTPoison.post(url, body, headers, http_options()) do
      {:ok, %HTTPoison.Response{status_code: 200, body: body}} ->
        response = Poison.decode!(body)
        {:ok, "Bearer #{response["access_token"]}"}
      {:ok, %HTTPoison.Response{status_code: status, body: ""}} ->
        {:error, status}
      {:ok, %HTTPoison.Response{status_code: status, body: body}} ->
        {:error, "#{status} #{body}"}
      {:error, %HTTPoison.Error{reason: :nxdomain}} ->
        {:error, "nxdomain"}
      {:error, %HTTPoison.Error{reason: reason}} ->
        {:error, reason}
      _ ->
        {:error, "Something went wrong"}
    end
  end

  def auth_url do
    {:system, tenant_id} = Application.get_env(:keyvault, :azure_tenant_id);
    "https://login.microsoftonline.com/#{tenant_id}/oauth2/token"
  end

  defp auth_body do
    {:system, client_id} = Application.get_env(:keyvault, :azure_client_id)
    {:system, client_secret} = Application.get_env(:keyvault, :azure_client_secret)
    {:form, [grant_type: "client_credentials", client_id: client_id, client_secret: client_secret, resource: "https://vault.azure.net"]}
  end

  defp http_options do
    [ssl: [{:versions, [:'tlsv1.2']}]]
  end
end

