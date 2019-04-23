defmodule Graphql.Router do
  use Plug.Router

  plug Plug.Parsers,
    parsers: [:urlencoded, :multipart, :json, Absinthe.Plug.Parser],
    pass: ["*/*"],
    json_decoder: Poison

  plug(:match)
  plug(:dispatch)

  forward "/graphql",
    to: Absinthe.Plug,
    init_opts: [schema: Graphql.Schema]

  forward "/graphiql",
    to: Absinthe.Plug.GraphiQL,
    init_opts: [schema: Graphql.Schema]

  # match(_, do: send_resp(conn, 404, "Oops!"))
  match(_, do: {})
end
