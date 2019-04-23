defmodule Graphql.User do
  use Ecto.Schema

  @primary_key {:user_id, :integer, autogenerate: false, source: :UserId}
  schema "Users" do
    field :external_id, Tds.UUID, source: :ExternalId
    field :email_identifier, :string, source: :EmailIdentifier
    field :login_name, :string, source: :LoginName
    field :contact_email, :string, source: :ContactEmail
    field :display_name, :string, source: :DisplayName
    field :first_name, :string, source: :FirstName
    field :last_name, :string, source: :LastName
    field :phone_number, :string, source: :PhoneNumber
    field :entity_state, :string, source: :EntityState
    field :created_date, :utc_datetime, source: :CreatedDate
    field :created_by, :string, source: :CreatedBy
    field :last_modified_date, :utc_datetime, source: :LastModifiedDate
    field :last_modified_by, :string, source: :LastModifiedBy
  end
end
