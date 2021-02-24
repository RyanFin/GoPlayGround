package main

import "fmt"

/*
 * Complete the timeConversion function below.
 */
func timeConversion(s string) string {
    var ret string

    // If the last 2 chars in the string == PM
    if(string(s[len(s)-2:]) == "PM"){
        // strip last two chars
        var removeEndStr = s[0:len(s)-2]

        fmt.Println(s[0:2])
        // if first two chars are 12
        switch s[0:2]{
            case "12":
            ret = removeEndStr
            break;
            case "01":
            ret = "13" + ":" + s[3: len(s)-2]
            break;
            case "02":
            ret = "14" + ":" + s[3: len(s)-2]
            break;
            case "03":
            ret = "15" + ":" + s[3: len(s)-2]
            break;
            case "04":
            ret = "16" + ":" + s[3: len(s)-2]
            break;
            case "05":
            ret = "17" + ":" + s[3: len(s)-2]
            break;
            case "06":
            ret = "18" + ":" + s[3: len(s)-2]
            break;
            case "07":
            ret = "19" + ":" + s[3: len(s)-2]
            break;
            case "08":
            ret = "20" + ":" + s[3: len(s)-2]
            break;
            case "09":
            ret = "21" + ":" + s[3: len(s)-2]
            break;
            case "10":
            ret = "22" + ":" + s[3: len(s)-2]
            break;
            case "11":
            ret = "23" + ":" + s[3: len(s)-2]
            break;

        }
    }

    // If the last 2 chars in the string == AM
    if(string(s[len(s)-2:]) == "AM"){


        if(s[0:2] == "12"){
            ret = "00" + ":" + s[3: len(s)-2]
        } else{
            // strip last two chars
            ret = s[0:len(s)-2]
        }

    }


    return ret
}

func main() {
    var str string

    fmt.Scan(&str)

    fmt.Println(timeConversion(str))
}
