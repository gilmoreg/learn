defmodule PlugDemoTest do
  use ExUnit.Case
  doctest PlugDemo

  test "greets the world" do
    assert PlugDemo.hello() == :world
  end
end
