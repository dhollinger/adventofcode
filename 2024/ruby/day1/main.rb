def main
  file = File.open("./input.txt")
  data = file.readlines.map(&:chomp)
  file.close

  left_list = []
  right_list = []
  diff_list = []
  sim_list = []

  data.each do |line|
    d = line.split
    left_list.append(d[0].to_i)
    right_list.append(d[1].to_i)
  end

  left_list = left_list.sort
  right_list = right_list.sort

  left_list.each_with_index do |n, index|
    distance = 0
    if left_list[index] < right_list[index]
      distance = right_list[index] - left_list[index]
    else
      distance = left_list[index] - right_list[index]
    end
    diff_list.append(distance)
  end

  total_dist = 0

  diff_list.each do |num|
    total_dist = total_dist + num
  end

  puts("Total Distance: #{total_dist}\n")

  # Part Two

  count = 0

  left_list.each do |n|
    right_list.each do |i|
      if n == i
        count = count + 1
      end
    end
    sim_num = n * count
    sim_list.append(sim_num)
    count = 0
  end

  total_sim = 0
  sim_list.each do |num|
    total_sim = total_sim + num
  end

  puts("Total Similarity Value: #{total_sim}\n")
end

main
