defmodule Graphql do
  @moduledoc """
  Documentation for Graphql.
  """

  @doc """
  Hello world.

  ## Examples

      iex> Graphql.hello()
      :world

  """
  use Application
  require Logger
  import Supervisor.Spec

  def start(_type, _args) do
    children = [
      Plug.Adapters.Cowboy.child_spec(:http, Graphql.Router, [], port: 8080),
      supervisor(Graphql.Repo, [])
    ]

    Logger.info("Started application")

    opts = [strategy: :one_for_one, name: Graphql.Supervisor]
    Supervisor.start_link(children, opts)
  end
end
