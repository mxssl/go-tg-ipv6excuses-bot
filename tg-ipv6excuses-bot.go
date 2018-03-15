package main

import (
	"time"
	"log"
	"math/rand"
	"os"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	tgToken := os.Getenv("tgToken")
	b, err := tb.NewBot(tb.Settings{
		Token:  tgToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	var excuses = [...]string{
		"We've still got plenty of IPv4",
		"It's not [mature](https://www.ietf.org/rfc/rfc2460.txt) enough",
		"Vendor bugs",
		"IPv6 is a [security risk](https://tools.ietf.org/html/rfc6204#section-4.4)",
		"None of our customers want it",
		"It's too complicated",
		"NAT444 is [fine](https://tools.ietf.org/html/rfc7021#section-3)",
		"[AWS doesn't support it](https://aws.amazon.com/about-aws/whats-new/2016/10/ipv6-support-for-cloudfront-waf-and-s3-transfer-acceleration/)",
		"IPv6 addresses are too long to remember",
		"[Android](https://code.google.com/p/android/issues/detail?id=32621) doesn't support DHCPv6",
		"No one else has [deployed](https://www.google.com/intl/en/ipv6/statistics.html) it",
		"It's too hard to support",
		"I don't want to [expose my MAC address](https://tools.ietf.org/html/rfc4941)",
		"It'll make it [easier](https://www.reddit.com/r/LabourUK/comments/3hkblt/im_stella_creasy_ama/cu84dwc) for the government to track me",
		"IPv6 is [slower](https://www.youtube.com/watch?v=An7s25FSK0U) than IPv4",
		"Larger headers are [less efficient](http://ipv6.com/articles/general/Top-10-Features-that-make-IPv6-greater-than-IPv4-Part4.htm)",
		"Hex is hard",
		"There's no certification track",
		"We don't need that many addresses",
		"We can use [RFC6598](https://tools.ietf.org/html/rfc6598)",
		"Our vendor doesn't support it",
		"It's not supported by [Google Compute](https://cloud.google.com/compute/docs/networking#networks)",
		"There's no ROI on deploying IPv6",
		"[Azure](https://azure.microsoft.com/en-us/pricing/faq/) doesn't support it",
		"We would have to rewrite our entire application to support it",
		"IPv6 is just a fad",
		"IPv6 isn't supported by [OVH Cloud](https://www.ovh.co.uk/vps/vps-cloud.xml)",
		"Those stupid Privacy Extension addresses [keep changing](https://tools.ietf.org/html/rfc7217#section-3)",
		"Can't we just [buy](http://ipv4marketgroup.com/ipv4-pricing-in-a-post-arin-runout-world/) more IPv4 addresses?",
		"Github doesn't support IPv6",
		"Our Lawful Intercept doesn't support IPv6 yet",
		"We'll deploy IPv6 next financial year",
		"IPv6 isn't an [Internet Standard](https://blog.apnic.net/2016/04/18/declaring-ipv6-internet-standard/) yet",
		"End users don't care about IPv6",
		"What's IPv6?",
		"Our recursive DNS can't handle the extra load",
		"What do you mean I have to wrap an IP in square brackets?",
		"I don't want to lose the [security](http://www.internetsociety.org/deploy360/blog/2015/01/ipv6-security-myth-3-no-ipv6-nat-means-less-security/) provided by NAT",
		"We don't have a lab to test it",
		"My [transit provider](http://he.net/cgi-bin/ip_transit_quote) doesn't support IPv6",
		"We'll deploy IPv6 right after we deploy DNSSEC",
		"It'll break our [GeoIP](https://www.maxmind.com/en/ipv6-information-and-faq)",
		"Our business intelligence team can't even parse IPv4 logfiles",
		"Our [DDOS](https://www.cloudflare.com/ipv6/) [mitigation](https://docs.incapsula.com/Content/read-more/ipv6-support.htm) platform can't monitor IPv6",
		"We forgot to include IPv6 in our last RFP",
		"Did you mean IPTV?",
		"We have no roadmap for native IPv6 as we [rolled out 6RD](https://ape3000.com/ipv6/)",
		"Too many people have [broken IPv6 stacks](https://en.wikipedia.org/wiki/IPv6_brokenness_and_DNS_whitelisting)",
		"We have IPv6, but we just want to keep things simple",
		"Our [Dynamic DNS](https://dynv6.com/) doesn't support it",
		"IPv6 just isn't a priority",
                "[Junos RESTful API service does not operate over IPv6](https://books.google.ru/books?id=EFoPDAAAQBAJ&pg=PT154&dq=Junos+RESTful+API+service+does+not+operate+over+IPv6&hl=en&sa=X&ved=0ahUKEwje6-D3iO_ZAhUEQJoKHV1ABSgQ6AEIKTAA#v=onepage&q=Junos%20RESTful%20API%20service%20does%20not%20operate%20over%20IPv6&f=false)",
	}

	b.Handle("/ipv6", func(m *tb.Message) {
		excuse := excuses[rand.Intn(len(excuses))]
		b.Send(m.Chat, excuse, tb.ParseMode(tb.ModeMarkdown), tb.NoPreview)
	})

	b.Handle("/help", func(m *tb.Message){
		b.Send(m.Chat, "Send me /ipv6 and get an excellent IPv6 excuse.")
	})

	b.Handle("/start", func(m *tb.Message){
		b.Send(m.Chat, "Hello! Send /help for help.")
	})

	b.Start()
}
