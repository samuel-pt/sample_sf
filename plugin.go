package sample_sf

import (
    "log"

    "github.com/itsabot/abot/shared/datatypes"
    "github.com/itsabot/abot/shared/nlp"
    "github.com/itsabot/abot/shared/plugin"
)

var p *dt.Plugin

func init() {
    trigger := &nlp.StructuredInput{
        Commands: []string{"find", "show", "get"},
        Objects: []string{"account", "opportunity", "contact"},
    }

    // Tell Abot how this plugin will respond to new conversations and follow-up
    // requests.
    fns := &dt.PluginFns{Run: Run, FollowUp: FollowUp}

    // Create the plugin.
    var err error
    pluginPath := "github.com/samuel-pt/sample_sf"
    p, err = plugin.New(pluginPath, trigger, fns)
    if err != nil {
        log.Fatalln("building", err)
    }
}

// Abot calls Run the first time a user interacts with a plugin
func Run(in *dt.Msg) (string, error) {
    return FollowUp(in)
}

// Abot calls FollowUp every subsequent time a user interacts with the plugin
// as long as the messages hit this plugin consecutively. As soon as Abot sends
// a message for this user to a different plugin, this plugin's Run function
// will be called the next it's triggered.  This Run/FollowUp design allows us
// to reset a plugin's state when a user changes conversations.
func FollowUp(in *dt.Msg) (string, error) {
    ret = fmt.Sprintf("Oh dear. You hit me to access salesforce ")
    return ret
}
