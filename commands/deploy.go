import (
  "os"
  "log"

  "github.com/charleswisniewski/hecate/client/operations"
  "github.com/go-openapi/strfmt"
  "github.com/go-openapi/spec"

  apiclient "github.com/charleswisniewski/hecate/client"
  httptransport "github.com/go-openapi/runtime/client"
)

func main() {


}
var cmd{$endpoint} = &cobra.Command{
    Use:   "${endpoint} -- Params",
    Short: "TODO: pull comments from model",
    Long: `TODO pull ong comments from model, maybe type information?`,
    Run: func(cmd *cobra.Command, args []string) {
      doRequest(args)
    },
}

d := models.Deploy{}
//  fmt.Printf("%+v\n", d)
  dr := reflect.ValueOf(&d).Elem()
  typeOfT := dr.Type()
  for i := 0; i < dr.NumField(); i++ {
    f := dr.Field(i)
    fmt.Printf("%s %s = %v\n", typeOfT.Field(i).Name, f.Type())//, f.Interface())
  //@TODO pull comments to feed in the description fields
    if (f.Type() == "int64") {
      varType = "IntVar"
    }  else {
      varType = "StringVar"
    }
    cmd${endpoint}.Flags().${varType}(&echoTimes, $key, "t", 1, "times to echo the input")
    var rootCmd = &cobra.Command{Use: "app"}
    rootCmd.AddCommand(cmd{$endpoint}, cmdEcho)
    //cmdEcho.AddCommand(cmdTimes)
    rootCmd.Execute()
}

func doRequest(deployArgs Deploy) {
  transport := httptransport.New(os.Getenv("TODOLIST_HOST"), "", nil)

  // create the API client, with the transport
  client := apiclient.New(transport, strfmt.Default)

  // to override the host for the default client
  // apiclient.Default.SetTransport(transport)

  // make the request to get all items
  //@TODO: iterate over model to create CLI asking for 

  resp, err := client.Deploy.PostDeploy(deployArgs)
 

  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%#v\n", resp.Payload)
}