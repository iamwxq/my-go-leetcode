package easy

import "fmt"

func ExcelSheetColumnTitleCase() {
    var columnNumber int = 2147483647
    column := excelSheetColumnTitleS1(columnNumber)
    fmt.Println(column)
}

func excelSheetColumnTitleS1(columnNumber int) string {
    var ans string
    var radix int = 26 // 26进制
    var base int = 64  // columnNumber是从1开始的 64 + 1 = 65 <=> 'A'

    if columnNumber <= radix {
        return fmt.Sprintf("%c", columnNumber+base)
    }

    remainder := columnNumber % radix
    ans += fmt.Sprintf("%c", remainder+base)
    columnNumber -= remainder

    for columnNumber != 0 {
        quotient := columnNumber / radix
        ans = fmt.Sprintf("%c", quotient+base) + ans
        columnNumber -= radix * quotient
    }

    return ans
}
