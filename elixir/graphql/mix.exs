defmodule Graphql.MixProject do
  use Mix.Project

  def project do
    [
      app: :graphql,
      version: "0.1.0",
      elixir: "~> 1.7",
      start_permanent: Mix.env() == :prod,
      deps: deps()
    ]
  end

  # Run "mix help compile.app" to learn about applications.
  def application do
    [
      mod: {Graphql, []},
      extra_applications: [:logger, :tds_ecto, :ecto],
    ]
  end

  # Run "mix help deps" to learn about dependencies.
  defp deps do
    [
      {:absinthe, "~> 1.4"},
      {:absinthe_ecto, ">= 0.0.0"},
      {:absinthe_plug, "~> 1.4"},
      {:cowboy, "~> 1.1.2"},
      {:tds_ecto, "~> 2.2"},
      {:poison, "~> 3.1"},
    ]
  end

  # Aliases are shortcuts or tasks specific to the current project.
  # For example, to create, migrate and run the seeds file at once:
  #
  #     $ mix ecto.setup
  #
  # See the documentation for `Mix` for more info on aliases.
  defp aliases do
    [
      "ecto.setup": ["ecto.create", "ecto.migrate", "run priv/repo/seeds.exs"],
      "ecto.reset": ["ecto.drop", "ecto.setup"],
      test: ["ecto.create --quiet", "ecto.migrate", "test"]
    ]
  end
end
