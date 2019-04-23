filename = IO.gets("File to count the words from: ") |> String.trim
words =
  File.read!(filename)
  |> String.split(~r{(\\n|[^\w'])+})
  |> Enum.filter(fn w -> w != "" end)
words |> Enum.count |> IO.puts
