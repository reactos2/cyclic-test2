package modDatabase


func CDBInitialize(dsn string) error {
  err := g_SingleGiga.Initialize(dsn)

  return err
}