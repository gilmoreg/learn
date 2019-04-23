defmodule Graphql.Schema.ContentTypes do
  use Absinthe.Schema.Notation

  object :user do
    field :external_id, :string
    field :user_id, :integer
    field :email_identifier, :string
    field :login_name, :string
    field :contact_email, :string
    field :display_name, :string
    field :first_name, :string
    field :last_name, :string
    field :phone_number, :string
    field :entity_state, :string
    field :created_date, :string
    field :created_by, :string
    field :last_modified_date, :string
    field :last_modified_by, :string
  end

  object :customer do
    field :customerId, :string
    field :externalId, :string
    field :displayName, :string
    field :primaryContactEmailAddress, :string
    field :primaryContactFirstName, :string
    field :primaryContactLastName, :string
    field :entityState, :string
  end
end
