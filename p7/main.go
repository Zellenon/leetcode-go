package p7

func reverse(x int) int {
	neg := x < 0;
	val := 0;

    if neg {
      if x < -9223372036854775808 {
        return 0;
		}

      x *= -1;
    }

    for x > 0 {
      if val < 214748365 {
        val *= 10;
      } else {
        return 0;
      }
      val += x % 10;
      x /= 10;
    }
    if (neg) {
      val *= -1;
    }
    return val;

}
