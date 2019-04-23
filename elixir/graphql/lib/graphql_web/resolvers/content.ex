defmodule Graphql.Resolvers.Content do
  def list_users(_parent, _args, _resolution) do
    {:ok, Graphql.Content.list_users()}
  end

  def get_user(_parent, %{id: id}, _resolution) do
    case Graphql.Content.find_user(id) do
      nil ->
        {:error, "User ID #{id} not found"}
      user ->
        {:ok, user}
    end
  end
end
