package cmd

import vultr "github.com/JamesClonk/vultr/lib"

func GetClient() *vultr.Client {
	return vultr.NewClient(*apiKey, nil)
}
