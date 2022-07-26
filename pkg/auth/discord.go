package auth

import (
	"golang.org/x/oauth2"
)

// All scope constants that can be used.
const (
	DiscordScopeIdentify                   = "identify"
	DiscordScopeBot                        = "bot"
	DiscordScopeEmail                      = "email"
	DiscordScopeGuilds                     = "guilds"
	DiscordScopeGuildsJoin                 = "guilds.join"
	DiscordScopeConnections                = "connections"
	DiscordScopeGroupDMJoin                = "gdm.join"
	DiscordScopeMessagesRead               = "messages.read"
	DiscordScopeRPC                        = "rpc"                    // Whitelist only
	DiscordScopeRPCAPI                     = "rpc.api"                // Whitelist only
	DiscordScopeRPCNotificationsRead       = "rpc.notifications.read" // Whitelist only
	DiscordScopeWebhookIncoming            = "webhook.Incoming"
	DiscordScopeApplicationsBuildsUpload   = "applications.builds.upload" // Whitelist only
	DiscordScopeApplicationsBuildsRead     = "applications.builds.read"
	DiscordScopeApplicationsStoreUpdate    = "applications.store.update"
	DiscordScopeApplicationsEntitlements   = "applications.entitlements"
	DiscordScopeRelationshipsRead          = "relationships.read" // Whitelist only
	DiscordScopeActivitiesRead             = "activities.read"    // Whitelist only
	DiscordScopeActivitiesWrite            = "activities.write"   // Whitelist only
	DiscordScopeApplicationsCommands       = "applications.command"
	DiscordScopeApplicationsCommandsUpdate = "applications.command.update"
)

// DiscordEndpoint is Discord's OAuth 2.0 endpoint.
var DiscordEndpoint = oauth2.Endpoint{
	AuthURL:   "https://discord.com/api/oauth2/authorize",
	TokenURL:  "https://discord.com/api/oauth2/token",
	AuthStyle: oauth2.AuthStyleInParams,
}
