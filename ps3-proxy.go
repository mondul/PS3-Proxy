package main

import (
    "github.com/elazarl/goproxy"
    "text/template"
    "net/http"
    "regexp"
    "bytes"
    "log"
    "net"
)

func ExternalIP() (string) {
    ifaces, err := net.Interfaces()

    if err != nil {
        return "localhost"
    }

    for _, iface := range ifaces {
        if iface.Flags&net.FlagUp == 0 {
            continue // interface down
        }

        if iface.Flags&net.FlagLoopback != 0 {
            continue // loopback interface
        }

        addrs, err := iface.Addrs()

        if err != nil {
            return "localhost"
        }

        for _, addr := range addrs {
            var ip net.IP

            switch v := addr.(type) {
                case *net.IPNet:
                    ip = v.IP

                case *net.IPAddr:
                    ip = v.IP
            }

            if ip == nil || ip.IsLoopback() {
                continue
            }

            ip = ip.To4()

            if ip == nil {
                continue // not an ipv4 address
            }

            return ip.String()
        }
    }

    return "localhost"
}

func main() {
    listTmpl := template.Must(template.New("list").Parse("# {{.Region}}\r\nDest={{.Dest}};ImageVersion=00000000;SystemSoftwareVersion=0.00;CDN=http://d{{.Code}}01.ps3.update.playstation.net/update/ps3/image/{{.Code}}/nodata;CDN_Timeout=30;"))

    type Region struct { Region, Dest, Code string }

    regions := map[string]Region{
        "jp": { "JP", "83", "jp" },
        "us": { "US", "84", "us" },
        "eu": { "EU", "85", "eu" },
        "kr": { "KR", "86", "kr" },
        "uk": { "UK", "87", "uk" },
        "mx": { "MX", "88", "mx" },
        "au": { "AU/NZ", "89", "au" },
        "sa": { "SouthAsia", "8A", "sa" },
        "tw": { "TW", "8B", "tw" },
        "ru": { "RU", "8C", "ru" },
        "cn": { "CN", "8D", "cn" },
    }

    proxy := goproxy.NewProxyHttpServer()

    proxy.OnRequest(goproxy.UrlMatches(regexp.MustCompile(`fus01\.ps3\.update\.playstation\.net\/update\/ps3\/list\/..\/ps3-updatelist\.txt`))).DoFunc(
        func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
            log.Println("[*] PS3 update query blocked")

            region, ok := regions[r.URL.Path[17:19]]

            if ok {
                buf := &bytes.Buffer{}
                listTmpl.Execute(buf, region)
                return r, goproxy.NewResponse(r,
                    goproxy.ContentTypeText, http.StatusOK,
                    buf.String())
            }

            return r, goproxy.NewResponse(r,
                goproxy.ContentTypeText, http.StatusNotFound,
                "File not found")
        })

    log.Println("[*] Starting PS3 proxy at " + ExternalIP() + ":8080 ...")
    http.ListenAndServe(":8080", proxy)
}
