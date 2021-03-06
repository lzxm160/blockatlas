package blockatlas

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_parseSubscriptions(t *testing.T) {
	tests := []struct {
		name          string
		subscriptions Webhook
		wantSubs      []Subscription
	}{
		{
			name: "webhook with 1 coin",
			subscriptions: Webhook{
				Webhook: "http://127.0.0.1:8080",
				Subscriptions: Subscriptions{
					"0": {"xpub6BpYi6J1GZzfY3yY7DbhLLccF3efQa18nQngM3jaehgtNSoEgk6UtPULpC3oK5oA3trczY8Ld34LFw1USMPfGHwTEizdD5QyGcMyuh2UoBA", "xpub6CYwPfnPJLPquufPkb98coSb3mdy1CgaZrWUtYWGJTJ4VWZUbzH9HLGy7nHpP7DG4UdTkYYpirkTWQSP7pWHsrk24Nos5oYNHpfr4BgPVTL"},
				},
			},
			wantSubs: []Subscription{
				{
					Coin: 0, Webhook: "http://127.0.0.1:8080",
					Address: "xpub6BpYi6J1GZzfY3yY7DbhLLccF3efQa18nQngM3jaehgtNSoEgk6UtPULpC3oK5oA3trczY8Ld34LFw1USMPfGHwTEizdD5QyGcMyuh2UoBA",
				},
				{
					Coin: 0, Webhook: "http://127.0.0.1:8080",
					Address: "xpub6CYwPfnPJLPquufPkb98coSb3mdy1CgaZrWUtYWGJTJ4VWZUbzH9HLGy7nHpP7DG4UdTkYYpirkTWQSP7pWHsrk24Nos5oYNHpfr4BgPVTL",
				},
			},
		},
		{
			name: "webhook with 2 coins",
			subscriptions: Webhook{
				Webhook: "http://127.0.0.1:8080",
				Subscriptions: Subscriptions{
					"2": {"zpub6rH4MwgyTmuexAX6HAraks5cKv5BbtmwdLirvnU5845ovUJb4abgjt9DtXK4ZEaToRrNj8dQznuLC6Nka4eMviGMinCVMUxKLpuyddcG9Vc"},
					"0": {"xpub6BpYi6J1GZzfY3yY7DbhLLccF3efQa18nQngM3jaehgtNSoEgk6UtPULpC3oK5oA3trczY8Ld34LFw1USMPfGHwTEizdD5QyGcMyuh2UoBA", "xpub6CYwPfnPJLPquufPkb98coSb3mdy1CgaZrWUtYWGJTJ4VWZUbzH9HLGy7nHpP7DG4UdTkYYpirkTWQSP7pWHsrk24Nos5oYNHpfr4BgPVTL"},
				},
			},
			wantSubs: []Subscription{
				{
					Coin: 2, Webhook: "http://127.0.0.1:8080",
					Address: "zpub6rH4MwgyTmuexAX6HAraks5cKv5BbtmwdLirvnU5845ovUJb4abgjt9DtXK4ZEaToRrNj8dQznuLC6Nka4eMviGMinCVMUxKLpuyddcG9Vc",
				},
				{
					Coin: 0, Webhook: "http://127.0.0.1:8080",
					Address: "xpub6BpYi6J1GZzfY3yY7DbhLLccF3efQa18nQngM3jaehgtNSoEgk6UtPULpC3oK5oA3trczY8Ld34LFw1USMPfGHwTEizdD5QyGcMyuh2UoBA",
				},
				{
					Coin: 0, Webhook: "http://127.0.0.1:8080",
					Address: "xpub6CYwPfnPJLPquufPkb98coSb3mdy1CgaZrWUtYWGJTJ4VWZUbzH9HLGy7nHpP7DG4UdTkYYpirkTWQSP7pWHsrk24Nos5oYNHpfr4BgPVTL",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSubs := tt.subscriptions.ParseSubscriptions()
			sort.Slice(gotSubs, func(i, j int) bool {
				return gotSubs[i].Coin > gotSubs[j].Coin
			})
			sort.Slice(tt.wantSubs, func(i, j int) bool {
				return tt.wantSubs[i].Coin > tt.wantSubs[j].Coin
			})
			assert.EqualValues(t, tt.wantSubs, gotSubs)
		})
	}
}
