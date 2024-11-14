package dataprocessor

type FundDataRecord struct {
    Date string
    Sum float64
}

type SingleFundData struct {
    Header string
    Records []FundDataRecord
    Sum float64
}

type FundData struct {
    Sums []SingleFundData
    TotalSum float64
}

func (fund *FundData) CalcSums() {
    var totalSum float64 = 0
    for i := range fund.Sums {
        var sum float64 = 0;
        for j:= range fund.Sums[i].Records {
            sum += fund.Sums[i].Records[j].Sum;
        }
        fund.Sums[i].Sum = sum;
        totalSum += sum
    }
    fund.TotalSum = totalSum
}
