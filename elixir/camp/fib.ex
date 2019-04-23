defmodule Fib do
  def time(func, arg) do
    t0 = Time.utc_now
    func.(arg)
    Time.diff(Time.utc_now, t0, :millisecond)
  end

  def compare(n \\ 40) do
    IO.puts "Naive: #{time(&naive/1, n)}ms"
    IO.puts "Faster: #{time(&faster/1, n)}ms"
  end

  def naive(1), do: 1
  def naive(2), do: 1
  def naive(n) do
    naive(n - 2) + naive(n - 1)
  end

  def faster(n), do: faster(n, 0, 1)
  def faster(1, _, acc2), do: acc2
  def faster(n, acc1, acc2) do
    faster(n - 1, acc2, acc1 + acc2)
  end
end
