defmodule MinimalTodo do
  def start() do
    get_command(nil)
  end

  def load_csv() do
    filename = IO.gets("Enter a .CSV file: ") |> String.trim
    read(filename)
      |> parse
      |> get_command
  end

  def save_csv() do

  end

  def get_command(data) do
    prompt = """
      Type the first letter of the command you want to run
      R)ead Todos  A)dd a Todo  D)elete a Todo    L)oad .CSV    S)ave .CSV
    """
    command = IO.gets(prompt)
      |> String.trim
      |> String.downcase

    case command do
      "r" -> show_todos(data)
      "a" -> add_todo(data)
      "l" -> load_csv()
      "s" -> save_csv()
      "d" -> delete_todo(data)
      "q" -> "Goodbye!"
      _   -> get_command(data)
    end
  end

  def add_todo(data) do
    name = get_item_name(data)
    fields = get_fields(data)
    # get_command(new_data)
  end

  def delete_todo(data) do
    todo = IO.gets("Which todo would you like to delete?\n") |> String.trim
    if Map.has_key? data, todo do
      IO.puts "ok."
      new_map = Map.drop(data, [todo])
      IO.puts ~s{#{todo} has been deleted.}
      get_command(new_map)
    else
      IO.puts ~s{There is no Todo named #{todo}!}
      show_todos(data, false)
      delete_todo(data)
    end
  end

  def read(filename) do
    case File.read(filename) do
      {:ok, body} -> body
      {:error, reason} -> IO.puts ~s(Could not open file "#{filename}"\n)
                          IO.puts ~s("#{:file.format_error reason}"\n)
                          get_command(nil)
    end
  end

  def parse(body) do
    [header | lines] = String.split(body, ~r{(\r\n|\r|\n)}) |> Enum.filter(fn x -> x != "" end)
    titles = tl String.split(header, ",")
    parse_lines(lines, titles)
  end

  def parse_lines(lines, titles) do
    Enum.reduce(lines, %{}, fn line, built ->
      [name | fields] = String.split(line, ",")
      if Enum.count(fields) == Enum.count(titles) do
        line_data = Enum.zip(titles, fields) |> Enum.into(%{})
        Map.merge(built, %{name => line_data})
      else
        built
      end
    end)
  end

  def get_item_name(data) do
    name = (IO.gets "Name: ") |> String.trim
    if Map.has_key?(data, name) do
      IO.puts "Todo with that name already exists"
      get_item_name(data)
    else
      name
    end
  end

  def get_fields(data) do
    data[hd Map.keys data] |> Map.keys
  end

  def show_todos(data, next? \\ true) do
    items = Map.keys data
    IO.puts "You have the following Todos:\n"
    Enum.each items, fn item -> IO.puts item end
    if next? do
      get_command(data)
    end
  end
end
