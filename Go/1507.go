package main

func reformatDate(date string) string {
    parts := strings.Split(date, " ")
    first := []byte(parts[0])
    first = first[:len(first)-2]
    if len(first) == 1 {
        first = []byte{'0', first[0]}
    }

    m := map[string]string{"Jan":"01", "Feb":"02", "Mar":"03", "Apr":"04", "May":"05", "Jun":"06", "Jul":"07", "Aug":"08", "Sep":"09", "Oct":"10", "Nov":"11", "Dec":"12"}
    second := m[parts[1]]
    return parts[2] + "-" + second + "-" + string(first)
}