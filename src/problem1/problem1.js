var sum_to_n_a = function(n) {
  // use loop to calculate the sum from 1 to n 
  var sum = 0;
  for (var i = 1; i <= n; i++) {
    sum += i;
  }
  return sum;
};

var sum_to_n_b = function(n) {
  // use mathematical formula to calculate the sum from 1 to n
  return n * (n + 1) / 2;
};

var sum_to_n_c = function(n) {
  // use recursion to calculate the sum from 1 to n
  if (n === 0) {
    return 0;
  }
  return n + sum_to_n_c(n - 1);
};