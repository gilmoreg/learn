defmodule AlchemyMarkdown do
  def to_html(markdown) do
    markdown
    |> big
    |> small
    |> hrs
    |> earmark
  end

  def earmark(markdown) do
    Earmark.as_html!((markdown || ""), %Earmark.Options{smartypants: false})
  end

  def big(markdown) do
    Regex.replace(~r/\+\+(.*)\+\+/, markdown, "<big>\\1</big>")
  end

  def small(markdown) do
    Regex.replace(~r/\-\-(.*)\-\-/, markdown, "<small>\\1</small>")
  end

  def hrs(markdown) do
    Regex.replace(~r{(^|\r\n|\r|\n)([-*])( *\2 *)+\2}, markdown, "\\1<hr />")
  end
end
