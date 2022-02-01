package plugin

type Plugin interface {
	Register(config map[string]interface{}, r *gin.RouterGroup) error
	Name() string
}
