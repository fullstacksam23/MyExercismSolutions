package isbnverifier

func IsValidISBN(isbn string) bool {
    ans := 0
    count := 10
	for _, r := range isbn {
        if r != '-'{
            if r == 'X'{
                if count == 1{
                    ans += count*10
                } else {
                    return false
                }
                
            } else {
            	ans += count*int(r-'0')
            }
            count-=1
        }
        
    }
    if count != 0{
        return false
    }
    return ans%11 == 0
}
