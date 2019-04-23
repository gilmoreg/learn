defmodule Auth.MixProject do
  use Mix.Project

  def project do
    [
      app: :auth,
      version: "0.1.0",
      elixir: "~> 1.7",
      start_permanent: Mix.env() == :prod,
      deps: deps()
    ]
  end

  # Run "mix help compile.app" to learn about applications.
  def application do
    [
      extra_applications: [:logger],
      mod: {Auth, []}
    ]
  end

  # Run "mix help deps" to learn about dependencies.
  defp deps do
    [
      {:guardian, "~> 1.0"},
      {:cowboy, "~> 1.1.2"},
      {:plug, "~> 1.3.4"}
    ]
  end
end
