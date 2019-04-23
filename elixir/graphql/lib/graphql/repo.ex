defmodule Graphql.Repo do
  use Ecto.Repo,
    otp_app: :graphql,
    adapter: Tds.Ecto
end
