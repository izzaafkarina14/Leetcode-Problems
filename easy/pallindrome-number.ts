function isPalindrome(x: number): boolean {
    if (x < 0) {
        return false
    }

    let input = x
    let reverseDigit = 0
    while (x > 0) {
        let lastDigit = x%10
        reverseDigit = reverseDigit*10 + lastDigit
        x = Math.floor(x/10)
    }
    return input == reverseDigit
};