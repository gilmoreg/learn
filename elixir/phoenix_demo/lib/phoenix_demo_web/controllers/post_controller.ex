defmodule PhoenixDemoWeb.PostController do
  use PhoenixDemoWeb, :controller

  def show(conn, %{"id" => id}) do
    json conn, %{id: id}
  end
end