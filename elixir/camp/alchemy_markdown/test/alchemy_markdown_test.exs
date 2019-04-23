defmodule AlchemyMarkdownTest do
  use ExUnit.Case
  doctest AlchemyMarkdown

  test "italicizes" do
    assert AlchemyMarkdown.to_html("*hello* world") =~ "<em>hello</em> world"
  end

  test "bolds" do
    assert AlchemyMarkdown.to_html("**hello** world") =~ "<strong>hello</strong> world"
  end

  test "expands big tags" do
    assert AlchemyMarkdown.to_html("++hello++ world") =~ "<big>hello</big> world"
  end

  test "expands small tags" do
    assert AlchemyMarkdown.to_html("--hello-- world") =~ "<small>hello</small> world"
  end

  test "expands hr tags" do
    str1 = "Stuff over the line\n---"
    str2 = "Stuff over the line\n***"
    str3 = "Stuff over the line\n- - -"
    str4 = "Stuff over the line\n*   *   *"
    Enum.each([str1, str2, str3, str4], fn str ->
      assert AlchemyMarkdown.hrs(str) == "Stuff over the line\n<hr />"
    end)
  end

  test "does not expand hrs not at start of line" do
    str = "Stuff over the line ---"
    assert AlchemyMarkdown.hrs(str) =~ str
  end
end
