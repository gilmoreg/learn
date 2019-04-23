defmodule KeyvaultTest do
  use ExUnit.Case
  doctest Keyvault

  test "greets the world" do
    assert Keyvault.hello() == :world
  end
end
