distance=
  File.stream!("day1.txt", :line)
  |> Enum.map(fn line ->
    line
    |> String.trim()
    |> String.split("   ")
    |> then(fn [s1, s2] ->
      { String.to_integer(s1), String.to_integer(s2) }
    end)
  end)
  |> Enum.unzip()
  |> then(fn ({arr1, arr2}) ->
    {Enum.sort(arr1), Enum.sort(arr2)}
  end)
  |> then(fn {arr1, arr2} ->
    Enum.zip(arr1, arr2)
    |> Enum.map(fn {i1, i2} -> abs(i1-i2) end)
    |> Enum.sum()
  end)

IO.puts("distance: #{distance}")
