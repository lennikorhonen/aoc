defmodule Day1 do
  def split(str) do
    String.split(str, " ", trim: true)
  end
end

list = File.stream!("./inputs/day_1/input.txt")
|> Enum.map(&String.trim/1)
|> Enum.map(&Day1.split/1)
|> Enum.to_list()

# left = Enum.map(list, fn x -> Enum.at(x, 0) end)
# |> Enum.map(&String.to_integer/1)
#
# right = Enum.map(list, fn x -> Enum.at(x, 1) end)
# |> Enum.map(&String.to_integer/1)


{left, right} =
  Enum.reduce(list, {[], []}, fn [l, r], {left_acc, right_acc} -> 
    {[String.to_integer(l) | left_acc], [String.to_integer(r) | right_acc]}
  end)

# Part 1 specific
left_sorted = Enum.sort(left)
right_sorted = Enum.sort(right)

total = Enum.zip([left_sorted, right_sorted])
|> Enum.map(fn {a, b} -> abs(b - a) end)
|> Enum.sum()

IO.inspect total
# Part 1 specific ends

# Part 2
left
|> Enum.map(fn x -> x * Enum.count(right, fn y -> x == y end) end)
|> Enum.sum()
|> IO.inspect
# Part 2 ends
