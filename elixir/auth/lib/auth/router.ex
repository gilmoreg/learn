defmodule Auth.Router do
  use Plug.Router

  plug(:match)
  plug(:dispatch)

  plug Guardian.Plug.VerifyHeader, key: :impersonate
  plug Guardian.Plug.EnsureAuthenticated, key: :impersonate

  get("/", do: send_resp(conn, 200, "Welcome"))
  match(_, do: send_resp(conn, 404, "Oops!"))
end
