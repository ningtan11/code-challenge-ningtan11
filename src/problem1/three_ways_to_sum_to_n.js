// We assume that n > 0
const firstTerm = 1;
const difference = 1;

var sum_to_n_a = function(n) {
    // your code here
    if (n <= 0) {
        return 0;
    }
    const noOfTerms = n;
    const result = (noOfTerms / 2) * ((2 * firstTerm) + (noOfTerms - 1) * difference);
    return result;
};

var sum_to_n_b = function(n) {
    // your code here
    var result = 0;
    for (var i = 1; i <= n; i++) {
        result += i;
    }
    return result;
};

var sum_to_n_c = function(n) {
    // your code here
    if (n < 0) {
        return 0;
    }
    const noOfTerm = n;
    const lastTerm = n;
    const result = noOfTerm * (firstTerm + lastTerm) / 2;
    return result
};
