var sum_to_n_a = function(n) {
    // your code here
    if (n < 0) {
        return 0;
    }
    const noOfTerms = n + 1;
    const result = (noOfTerms / 2) * (noOfTerms - 1);
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
    const noOfTerm = n + 1;
    const lastTerm = n;
    const result = noOfTerm * (0 + lastTerm) / 2;
    return result
};