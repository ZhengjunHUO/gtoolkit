package gtoolkit

func GetCatalanNum(n int) int {
        rslt := 1

        // C(2n,n)/(n+1)  =>  (2n)! / (n)!(2n-n)! / (n+1)
        for i:=1; i<=n; i++ {
                rslt = rslt*(i+n)/i
        }
        
        return rslt/(n+1)
}
