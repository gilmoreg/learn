defmodule Graphql.Schema do
  require Logger
  use Absinthe.Schema
  import_types Graphql.Schema.ContentTypes

  alias Graphql.Resolvers

  query do
    @desc "Get all users"
    field :users, list_of(:user) do
      Logger.info("get all users query")
      resolve &Resolvers.Content.list_users/3
    end

    @desc "Get a user"
    field :user, :user do
      arg :id, non_null(:id)
      resolve &Resolvers.Content.get_user/3
    end

  end

end
