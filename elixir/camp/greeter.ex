defmodule Greeter do
  def start do
    name = IO.gets "Hi! What's your name?\n"
    case String.trim(name) do
      "Grayson" -> IO.puts "That's my favorite name!"
      name -> IO.puts "Hi there #{name}!"
    end
  end
end
