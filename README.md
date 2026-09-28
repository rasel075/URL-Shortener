# URL-Shortener
# Go URL Shortener — সহজ বাংলায় ব্যাখ্যা

এই প্রোগ্রামটি একটি **ছোট URL Shortener**।

যেমন:

```text
https://google.com
```

কে ছোট করে:

```text
http://localhost:8080/70061
```

তারপর `http://localhost:8080/70061` এ গেলে আবার `https://google.com`-এ নিয়ে যাবে।

---

# 1. Package

```go
package main
```

`package main` মানে এটি আমাদের main program।

Go program চালু হলে `main()` function থেকে execution শুরু হবে।

---

# 2. Import

```go
import (
    "fmt"
    "net/http"
)
```

এখানে আমরা ২টা package ব্যবহার করছি।

### `fmt`

```go
"fmt"
```

Screen-এ কিছু print করার জন্য ব্যবহার করি।

যেমন:

```go
fmt.Println("Hello")
```

---

### `net/http`

```go
"net/http"
```

HTTP server বানানোর জন্য ব্যবহার করছি।

এটার সাহায্যে আমরা:

* Server চালাব
* URL/request নেব
* Response দেব
* Redirect করব

---

# 3. URL রাখার Map

```go
var urls = map[string]string{}
```

এটা একটি **map**।

Map-এর কাজ হলো:

```text
Key → Value
```

এখানে:

```text
Short Code → Original URL
```

রাখব।

যেমন:

```text
70061 → https://google.com
```

তাই code-এ হবে:

```go
urls["70061"] = "https://google.com"
```

এখন:

```go
urls["70061"]
```

লিখলে আমরা পাব:

```text
https://google.com
```

---

# 4. `shorten()` Function

```go
func shorten(w http.ResponseWriter, r *http.Request) {
```

এই function-এর কাজ হলো **long URL-কে short URL বানানো**।

এখানে দুইটা জিনিস আছে:

```go
w
```

এটা দিয়ে আমরা user-কে response পাঠাব।

আর:

```go
r
```

এর মধ্যে user-এর request থাকে।

সহজভাবে:

```text
r = User কী চেয়েছে
w = User-কে কী উত্তর দেব
```

---

# 5. User-এর দেওয়া URL নেওয়া

```go
longURL := r.URL.Query().Get("url")
```

ধরো browser-এ আমরা লিখলাম:

```text
http://localhost:8080/shorten?url=https://google.com
```

এখানে:

```text
/shorten
```

হলো path।

আর:

```text
?url=https://google.com
```

হলো query parameter।

আমরা এই:

```text
https://google.com
```

অংশটা বের করতে লিখেছি:

```go
r.URL.Query().Get("url")
```

তাই:

```go
longURL
```

এর মধ্যে থাকবে:

```text
https://google.com
```

---

# 6. URL দেওয়া হয়েছে কিনা check

```go
if longURL == "" {
```

এর মানে:

> যদি `longURL` খালি হয়।

যেমন user যদি শুধু যায়:

```text
http://localhost:8080/shorten
```

তাহলে কোনো URL দেওয়া হয়নি।

তখন:

```go
longURL == ""
```

হবে।

---

```go
fmt.Fprintln(w, "URL is required")
```

User-কে বলবে:

```text
URL is required
```

অর্থাৎ:

> URL দিতে হবে।

---

```go
return
```

এর মানে:

> এখানেই function বন্ধ করো।

তাই নিচের code আর execute হবে না।

---

# 7. Short Code

```go
code := "70061"
```

এখানে আমরা short code হিসেবে:

```text
70061
```

নিয়েছি।

তাহলে:

```text
Long URL
https://google.com

        ↓

Short Code
70061
```

---

# 8. URL save করা

```go
urls[code] = longURL
```

ধরো:

```go
code = "70061"
```

এবং:

```go
longURL = "https://google.com"
```

তাহলে line-টা আসলে এমন হবে:

```go
urls["70061"] = "https://google.com"
```

মানে:

```text
70061 → https://google.com
```

এই সম্পর্কটা map-এর মধ্যে save হয়ে গেল।

---

# 9. Short URL দেখানো

```go
fmt.Fprintln(w, "Short URL: http://localhost:8080/"+code)
```

এখানে:

```go
code
```

হলো:

```text
70061
```

তাই user দেখবে:

```text
Short URL: http://localhost:8080/70061
```

---

# এখন পর্যন্ত কী হলো?

User লিখল:

```text
http://localhost:8080/shorten?url=https://google.com
```

Program:

```text
URL নিল
   ↓
https://google.com
   ↓
code নিল
   ↓
70061
   ↓
Map-এ রাখল

70061 → https://google.com
   ↓
User-কে short URL দিল

http://localhost:8080/70061
```

---

# 10. এবার `redirect()` Function

```go
func redirect(w http.ResponseWriter, r *http.Request) {
```

এই function-এর কাজ হলো:

> Short URL থেকে original URL বের করে user-কে সেখানে পাঠানো।

যেমন:

```text
http://localhost:8080/70061
```

এখান থেকে:

```text
https://google.com
```

বের করবে।

---

# 11. Short Code বের করা

```go
code := r.URL.Path[1:]
```

এটা একটু গুরুত্বপূর্ণ।

User যদি যায়:

```text
http://localhost:8080/70061
```

তাহলে:

```go
r.URL.Path
```

হবে:

```text
/70061
```

কিন্তু আমাদের দরকার:

```text
70061
```

`/` দরকার নেই।

তাই:

```go
[1:]
```

ব্যবহার করা হয়েছে।

এটা প্রথম character বাদ দেয়।

```text
/70061
^
প্রথম character
```

বাদ দিলে:

```text
70061
```

তাই:

```go
code = "70061"
```

---

# 12. Root URL Check

```go
if code == "" {
```

যদি user শুধু যায়:

```text
http://localhost:8080/
```

তাহলে:

```go
r.URL.Path
```

হবে:

```text
/
```

এবং:

```go
r.URL.Path[1:]
```

হবে:

```text
""
```

তাই `code` empty।

---

```go
fmt.Fprintln(w, "URL Shortener is running")
```

তখন user দেখবে:

```text
URL Shortener is running
```

মানে server ঠিকভাবে চলছে।

---

```go
return
```

তারপর function শেষ।

---

# 13. Map থেকে URL বের করা

```go
longURL := urls[code]
```

ধরো:

```go
code = "70061"
```

তাহলে:

```go
urls[code]
```

মানে:

```go
urls["70061"]
```

আমাদের map-এ ছিল:

```text
70061 → https://google.com
```

তাই:

```go
longURL
```

হয়ে যাবে:

```text
https://google.com
```

---

# 14. Short URL পাওয়া গেছে কিনা check

```go
if longURL == "" {
```

যদি map-এর মধ্যে code না থাকে, তাহলে `longURL` empty হবে।

যেমন user যদি যায়:

```text
http://localhost:8080/99999
```

কিন্তু:

```text
99999
```

আমাদের map-এ নেই।

তাহলে:

```go
longURL == ""
```

হবে।

---

```go
fmt.Fprintln(w, "Short URL not found")
```

তখন user দেখবে:

```text
Short URL not found
```

অর্থাৎ:

> এই short URL-এর কোনো original URL পাওয়া যায়নি।

---

```go
return
```

Function শেষ।

---

# 15. আসল Redirect

```go
http.Redirect(w, r, longURL, http.StatusFound)
```

এটাই সবচেয়ে গুরুত্বপূর্ণ line।

ধরো:

```go
longURL = "https://google.com"
```

তখন browser-কে বলা হবে:

> `https://google.com` এ চলে যাও।

ফলে:

```text
http://localhost:8080/70061
            ↓
      https://google.com
```

---

## `http.StatusFound` কী?

```go
http.StatusFound
```

এর মানে HTTP status code:

```text
302
```

এটা browser-কে বলে:

> অন্য একটা URL-এ যাও।

---

# 16. `main()` Function

```go
func main() {
```

Go program এখান থেকেই শুরু হয়।

---

# 17. `/shorten` Route তৈরি

```go
http.HandleFunc("/shorten", shorten)
```

এর মানে:

```text
/shorten
   ↓
shorten function
```

অর্থাৎ user যখন:

```text
http://localhost:8080/shorten?url=https://google.com
```

এ যাবে, তখন:

```go
shorten()
```

function চলবে।

---

# 18. `/` Route

```go
http.HandleFunc("/", redirect)
```

এর মানে হলো `/` দিয়ে শুরু হওয়া request-এর জন্য `redirect()` function ব্যবহার হবে।

যেমন:

```text
/70061
/12345
/abcde
```

এসব request `redirect()` function-এ যাবে।

---

# 19. Server-এর Message

```go
fmt.Println("Server running on port 8080")
```

Terminal-এ দেখাবে:

```text
Server running on port 8080
```

মানে:

> Server চালু হয়েছে।

---

# 20. Server Start

```go
http.ListenAndServe(":8080", nil)
```

এটা আসলে server চালু করে।

```text
:8080
```

মানে server **8080 port**-এ চলবে।

তাই আমরা browser থেকে:

```text
http://localhost:8080
```

দিয়ে access করতে পারব।

---

# পুরো Program একসাথে বুঝি

ধরো আমরা প্রথমে যাই:

```text
http://localhost:8080/shorten?url=https://google.com
```

### Step 1

```go
longURL := r.URL.Query().Get("url")
```

পাবে:

```text
https://google.com
```

### Step 2

```go
code := "70061"
```

পাবে:

```text
70061
```

### Step 3

```go
urls[code] = longURL
```

Map-এ save হবে:

```text
70061 → https://google.com
```

### Step 4

User-কে দেওয়া হবে:

```text
http://localhost:8080/70061
```

---

## এরপর user short URL-এ গেলে

```text
http://localhost:8080/70061
```

### Step 5

```go
code := r.URL.Path[1:]
```

পাবে:

```text
70061
```

### Step 6

```go
longURL := urls[code]
```

Map থেকে পাবে:

```text
https://google.com
```

### Step 7

```go
http.Redirect(w, r, longURL, http.StatusFound)
```

Browser চলে যাবে:

```text
https://google.com
```

---

# সবচেয়ে সহজভাবে মনে রাখার নিয়ম

এই project-এ মূলত **দুইটা কাজ** হচ্ছে।

### 1️⃣ Long URL → Short URL

```text
https://google.com
        ↓
     70061
        ↓
http://localhost:8080/70061
```

এটা করে:

```go
shorten()
```

---

### 2️⃣ Short URL → Long URL

```text
http://localhost:8080/70061
        ↓
      70061
        ↓
https://google.com
```

এটা করে:

```go
redirect()
```

---

# পুরো Code-এর দায়িত্ব

```go
var urls = map[string]string{}
```

👉 URL জমা রাখে।

```go
shorten()
```

👉 Long URL নিয়ে short URL বানায়।

```go
redirect()
```

👉 Short URL থেকে original URL বের করে।

```go
http.HandleFunc()
```

👉 কোন URL-এ কোন function চলবে সেটা ঠিক করে।

```go
http.ListenAndServe()
```

👉 Server চালু করে।

```go
http.Redirect()
```

👉 User-কে অন্য URL-এ পাঠায়।

---

# একটা গুরুত্বপূর্ণ সমস্যা

এই code-এ:

```go
code := "70061"
```

সবসময় একই code।

তাই:

```text
Google
   ↓
70061
```

এরপর:

```text
YouTube
   ↓
70061
```

দিলে আগেরটা replace হয়ে যাবে।

কারণ map-এ হবে:

```text
70061 → YouTube
```

তাই পরবর্তীতে `70061` খুললে Google নয়, YouTube-এ যাবে।

**বাস্তব URL shortener-এ প্রত্যেক URL-এর জন্য আলাদা unique code generate করতে হয়।**
