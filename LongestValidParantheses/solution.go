type ValidParantheses struct {
    start int
    end int
    appendedLength int
}

func longestValidParentheses(s string) int {
    var startingParanthesesPosStack []int
    var reducedValidParantheses []ValidParantheses
    result := 0
    currentLength := 0

    for p,c := range s {
        if c == '(' {
            // new parantheses sequence
            startingParanthesesPosStack = append(startingParanthesesPosStack, p)
        } else {
            l := len(startingParanthesesPosStack)
            if l > 0 {
                // valid parantheses closed
                currentStartingPos := startingParanthesesPosStack[len(startingParanthesesPosStack)-1]
                appendedLength := 0
                startingParanthesesPosStack = startingParanthesesPosStack[:len(startingParanthesesPosStack)-1]

                currentLength = p - currentStartingPos + 1
                // check if the previous detected valid parentheses is inside this one (there can only be one)
                if len(reducedValidParantheses) > 0 && reducedValidParantheses[len(reducedValidParantheses)-1].start > currentStartingPos {
                    // remove last element
                    reducedValidParantheses = reducedValidParantheses[:len(reducedValidParantheses)-1]
                }

                // check if the previous valid parantheses ends just before the current one starts (there can only be one)
                if len(reducedValidParantheses) > 0 && reducedValidParantheses[len(reducedValidParantheses)-1].end + 1 == currentStartingPos {
                    // pop last element
                    appendedParantheses := reducedValidParantheses[len(reducedValidParantheses)-1]
                    reducedValidParantheses = reducedValidParantheses[:len(reducedValidParantheses)-1]

                    appendedLength = appendedParantheses.appendedLength + appendedParantheses.end - appendedParantheses. start + 1
                    currentLength += appendedLength
                }

                if currentLength > result {
                    result = currentLength
                }
                reducedValidParantheses = append(reducedValidParantheses, ValidParantheses{currentStartingPos, p, appendedLength})
            }
        }
    }
    return result
}
