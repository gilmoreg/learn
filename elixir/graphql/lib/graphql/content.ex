defmodule Graphql.Content do
  alias Graphql.{Repo, User}

  def list_users() do
    Repo.all(User)
  end

  def find_user(id) do
    Repo.get(User, id)
  end
end
