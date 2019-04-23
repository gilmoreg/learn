defmodule StatWatch.Scheduler do
  use GenServer

  def start_link do
    GenServer.start_link(__MODULE__, %{})
  end

  def init(state) do
    handle_info(:work, state)
    {:ok, state}
  end

  def handle_info(:work, state) do
    StatWatch.fetch_stats()
    schedule_work()
    {:noreply, state}
  end

  defp schedule_work() do
    # Repeat every 6 hours
    Process.send_after(self(), :work, 6 * 60 * 60 * 1000)
  end
end
