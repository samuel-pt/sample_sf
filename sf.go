package sample_sf

import (
	"log"

	"github.com/itsabot/abot/shared/datatypes"
	"github.com/itsabot/abot/shared/nlp"
	"github.com/itsabot/abot/shared/plugin"
)

var p *dt.Plugin

func init() {
	// Creating new Plugin
	// This is the first step to be done
	var err error
	p, err = plugin.New("github.com/samuel-pt/sample_sf")
	if err != nil {
		log.Fatal(err)
	}

	// Keywords to be used by this plugin
	// Here two functions are registered, handleSF and handleWave
	// Objects "sf", "object", "opportunity"
	plugin.SetKeywords(p,
		dt.KeywordHandler{
			Fn: handleSF,
			Trigger: &nlp.StructuredInput{
				Commands: []string{"what", "show", "tell",
					"how", "is"},
				Objects: []string{"sf", "object",
					"opportunity", "account", "lead"},
			},
		},
		dt.KeywordHandler{
			Fn: handleWave,
			Trigger: &nlp.StructuredInput{
				Commands: []string{"what", "show", "tell",
					"how", "is"},
				Objects: []string{"wave", "analytics", "dataset",
					"dashboard", "lens"},
			},
		},
	)
	plugin.SetStates(p, [][]dt.State{[]dt.State{
		dt.State{
			OnEntry: func(in *dt.Msg) string {
				return "What is your salesforce username?"
			},
			OnInput: func(in *dt.Msg) {
				if len(cities) > 0 {
					p.SetMemory(in, "sf_username", in.Sentence)
				}
			},
			Complete: func(in *dt.Msg) (bool, string) {
				return p.HasMemory(in, "sf_username"), ""
			},
			SkipIfComplete: true,
		},
		dt.State{
			OnEntry: func(in *dt.Msg) string {
				return "What is your salesforce password?"
			},
			OnInput: func(in *dt.Msg) {
				if len(cities) > 0 {
					p.SetMemory(in, "sf_password", "*******")
				}
			},
			Complete: func(in *dt.Msg) (bool, string) {
				return p.HasMemory(in, "sf_password"), ""
			},
			SkipIfComplete: true,
		},
	}})

	if err = plugin.Register(p); err != nil {
		p.Log.Fatal(err)
	}
}

func handleSF(in *dt.Msg) (resp string) {
	sf_username := p.GetMemory(in, "sf_username")
	p.Log.Info(fmt.Sprintf("SF User from Cache %s  ", sf_username))
	sf_password := p.GetMemory(in, "sf_password")
	p.Log.Info(fmt.Sprintf("SF Password from Cache %s  ", sf_password))

	return fmt.Springf("Haha... Trying to access SF with %s and with password %s",
		sf_username, sf_password)
}

func handleWave(in *dt.Msg) (resp string) {
	sf_username := p.GetMemory(in, "sf_username")
	p.Log.Info(fmt.Sprintf("SF User from Cache %s  ", sf_username))
	sf_password := p.GetMemory(in, "sf_password")
	p.Log.Info(fmt.Sprintf("SF Password from Cache %s  ", sf_password))

	return fmt.Springf("Haha... Trying to access SF Wave with %s and with password %s",
		sf_username, sf_password)
}
