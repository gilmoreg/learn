defmodule PlugDemo do
  @moduledoc """
  Documentation for PlugDemo.
  """

  @doc """
  Hello world.

  ## Examples

      iex> PlugDemo.hello()
      :world

  """
  use Application
  require Logger

  def start(_type, _args) do
    children = [
      Plug.Adapters.Cowboy.child_spec(:http, PlugDemo.Router, [], port: 8080)
    ]

    Logger.info("Started application")
    Supervisor.start_link(children, strategy: :one_for_one)
  end
end
