package main

import (
	"fmt"
)

// تابعی برای تولید اعداد طبیعی و ارسال آن‌ها به کانال
func generateNumbers(ch chan int, limit int) {
	for i := 2; i <= limit; i++ {
		ch <- i
	}
	close(ch) // بستن کانال پس از ارسال تمام اعداد
}

// فیلتر کردن اعداد غیر اول و ارسال فقط اعداد اول به کانال out
func filterPrimes(in chan int, out chan int, prime int) {
	for num := range in {
		if num%prime != 0 {
			out <- num
		}
	}
	close(out) // بستن کانال پس از پایان پردازش
}

func main() {
	limit := 50
	ch := make(chan int) // کانال اصلی تولید اعداد

	go generateNumbers(ch, limit) // گوروتین تولید اعداد

	for {
		prime, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(prime, "is prime") // چاپ عدد اول

		chFiltered := make(chan int)           // کانال فیلتر شده برای اعداد بعدی
		go filterPrimes(ch, chFiltered, prime) // فیلتر اعداد غیر اول
		ch = chFiltered                        // ادامه فیلتر با کانال جدید
	}
}
