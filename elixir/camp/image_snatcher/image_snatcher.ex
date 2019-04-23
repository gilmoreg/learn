defmodule ImageSnatcher do
  def start() do
    get_cwd_list() |> sort_files
  end

  def get_cwd_list() do
    files = File.ls(".")
    case files do
      {:ok, files} -> files
      {:error, reason} -> IO.puts ~s{#{:file.format_error(reason)}}
    end
  end

  def sort_files(files) do
    Enum.each(files, fn file ->
      extension = String.split(file, ".") |> List.last
      case extension do
        "bmp" -> move_file(file)
        "jpg" -> move_file(file)
        "png" -> move_file(file)
        _ -> nil
      end
    end)
  end

  def move_file(file) do
    ensure_target_dir()
    case File.cp(file, "images/#{file}") do
      :ok -> case File.rm(file) do
        :ok -> :ok
        {:error, reason } -> IO.puts ~s{#{:file.format_error(reason)}}
      end
      {:error, reason } -> IO.puts ~s{#{:file.format_error(reason)}}
    end
  end

  def ensure_target_dir() do
    case File.mkdir("images") do
      :ok -> :ok
      {:error, :eexists } -> :ok
      {:error, reason } -> IO.puts ~s{#{:file.format_error(reason)}}
    end
  end
end
