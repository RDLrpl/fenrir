package accs

type Sessions struct {
	Telegram []TGSession `json:"Sessions"`
}

type TGSession struct {
	ComparedProxy   Prx `json:"proxy"`
	ComparedAccount Acc `json:"account"`
}

type Prx struct {
	Protocol  string `json:"protocol"`
	Transport string `json:"transport"`
	IpAddress string `json:"ip"`
	Port      string `json:"port"`
	Login     string `json:"login"`
	Pass      string `json:"pass"`
}

type Acc struct {
	ID          string `json:"visualname"`
	API_id      string `json:"api"`
	API_hash    string `json:"hash"`
	PhoneNumber string `json:"number"`
	CloudPass   string `json:"pass"`
}

func TGCompare(Accs Accounts, Proxs Proxies) Sessions {
	var Sessions Sessions

	for _, api := range Accs.TGApis {
		var session TGSession
		var prox Prx
		var acc Acc

		Target := api.Prox

		for _, proxy := range Proxs.Proxies {
			if proxy.Tid == Target {
				prox = Prx{
					Protocol:  proxy.Proto,
					Transport: proxy.Transport,
					IpAddress: proxy.Ip,
					Port:      proxy.Port,
					Login:     proxy.Login,
					Pass:      proxy.Pass,
				}
			}
		}

		acc = Acc{
			ID:          api.ID,
			API_id:      api.API_id,
			API_hash:    api.API_hash,
			PhoneNumber: api.Number,
			CloudPass:   api.Pass,
		}

		session = TGSession{
			ComparedProxy:   prox,
			ComparedAccount: acc,
		}

		Sessions.Telegram = append(Sessions.Telegram, session)
	}

	return Sessions
}
