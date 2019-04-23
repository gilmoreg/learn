defmodule Graphql.Customer do
  use Ecto.Schema

  schema "customers" do
    field :customerId, :string
    field :externalId, :string
    field :displayName, :string
    field :primaryContactEmailAddress, :string
    field :primaryContactFirstName, :string
    field :primaryContactLastName, :string
    field :entityState, :string
  end

end
