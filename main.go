package main

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

type response struct {
	Data struct {
		Data struct {
			SearchDashClustersByAll struct {
				Metadata struct {
					EntityResultAttributes  any      `json:"entityResultAttributes"`
					TotalResultCount        int      `json:"totalResultCount"`
					SecondaryFilterCluster  any      `json:"secondaryFilterCluster"`
					RecipeTypes             []string `json:"$recipeTypes"`
					LazyRightRail           any      `json:"lazyRightRail"`
					QueryType               any      `json:"queryType"`
					Type                    string   `json:"$type"`
					PrimaryResultType       string   `json:"primaryResultType"`
					PaginationToken         any      `json:"paginationToken"`
					PrimaryFilterCluster    any      `json:"primaryFilterCluster"`
					BlockedQuery            bool     `json:"blockedQuery"`
					EntityActionButtonStyle any      `json:"entityActionButtonStyle"`
					SearchID                string   `json:"searchId"`
					FilterAppliedCount      int      `json:"filterAppliedCount"`
					ClusterTitleFontSize    string   `json:"clusterTitleFontSize"`
					SimpleInsightAttributes any      `json:"simpleInsightAttributes"`
					KnowledgeCardRightRail  any      `json:"knowledgeCardRightRail"`
				} `json:"metadata"`
				Paging struct {
					Count       int      `json:"count"`
					Start       int      `json:"start"`
					Total       int      `json:"total"`
					RecipeTypes []string `json:"$recipeTypes"`
					Type        string   `json:"$type"`
				} `json:"paging"`
				RecipeTypes []string `json:"$recipeTypes"`
				Elements    []struct {
					Image                any      `json:"image"`
					QuickFilterActions   []any    `json:"quickFilterActions"`
					ClusterRenderType    string   `json:"clusterRenderType"`
					Dismissable          bool     `json:"dismissable"`
					TotalResultCount     any      `json:"totalResultCount"`
					ControlName          any      `json:"controlName"`
					Description          any      `json:"description"`
					Title                any      `json:"title"`
					RecipeTypes          []string `json:"$recipeTypes"`
					Type                 string   `json:"$type"`
					ActionTypeName       any      `json:"actionTypeName"`
					NavigationText       any      `json:"navigationText"`
					Feature              any      `json:"feature"`
					NavigationCardAction any      `json:"navigationCardAction"`
					Position             int      `json:"position"`
					Items                []struct {
						Item struct {
							EntityResult           any `json:"entityResult"`
							KeywordsSuggestionCard any `json:"keywordsSuggestionCard"`
							Cluster                any `json:"cluster"`
							SimpleText             struct {
								TextDirection                 string   `json:"textDirection"`
								Text                          string   `json:"text"`
								AttributesV2                  []any    `json:"attributesV2"`
								AccessibilityTextAttributesV2 []any    `json:"accessibilityTextAttributesV2"`
								AccessibilityText             any      `json:"accessibilityText"`
								RecipeTypes                   []string `json:"$recipeTypes"`
								Type                          string   `json:"$type"`
							} `json:"simpleText"`
							QueryClarificationCard any `json:"queryClarificationCard"`
							BannerCard             any `json:"bannerCard"`
							PromoCard              any `json:"promoCard"`
							CenteredText           any `json:"centeredText"`
							SearchSuggestionCard   any `json:"searchSuggestionCard"`
							SimpleImage            any `json:"simpleImage"`
							FeedbackCard           any `json:"feedbackCard"`
							KnowledgeCardV2        any `json:"knowledgeCardV2"`
						} `json:"item"`
						Position    int      `json:"position"`
						RecipeTypes []string `json:"$recipeTypes"`
						Type        string   `json:"$type"`
					} `json:"items"`
					Results    []any  `json:"results"`
					TrackingID string `json:"trackingId"`
				} `json:"elements"`
				Type string `json:"$type"`
			} `json:"searchDashClustersByAll"`
			RecipeTypes []string `json:"$recipeTypes"`
			Type        string   `json:"$type"`
		} `json:"data"`
	} `json:"data"`
	Included []struct {
		EntityUrn                string   `json:"entityUrn"`
		RecipeTypes              []string `json:"$recipeTypes"`
		Type                     string   `json:"$type"`
		Template                 string   `json:"template,omitempty"`
		ActorNavigationContext   any      `json:"actorNavigationContext,omitempty"`
		TrackingUrn              string   `json:"trackingUrn,omitempty"`
		ControlName              any      `json:"controlName,omitempty"`
		InterstitialComponent    any      `json:"interstitialComponent,omitempty"`
		PrimaryActions           []any    `json:"primaryActions,omitempty"`
		EntityCustomTrackingInfo struct {
			MemberDistance                 string   `json:"memberDistance"`
			PrivacySettingsInjectionHolder any      `json:"privacySettingsInjectionHolder"`
			RecipeTypes                    []string `json:"$recipeTypes"`
			NameMatch                      bool     `json:"nameMatch"`
			Type                           string   `json:"$type"`
		} `json:"entityCustomTrackingInfo,omitempty"`
		Title struct {
			TextDirection                 string   `json:"textDirection"`
			Text                          string   `json:"text"`
			AttributesV2                  []any    `json:"attributesV2"`
			AccessibilityTextAttributesV2 []any    `json:"accessibilityTextAttributesV2"`
			AccessibilityText             any      `json:"accessibilityText"`
			RecipeTypes                   []string `json:"$recipeTypes"`
			Type                          string   `json:"$type"`
		} `json:"title,omitempty"`
		OverflowActions           []any `json:"overflowActions,omitempty"`
		SearchActionType          any   `json:"searchActionType,omitempty"`
		ActorInsights             []any `json:"actorInsights,omitempty"`
		InsightsResolutionResults []any `json:"insightsResolutionResults,omitempty"`
		BadgeIcon                 any   `json:"badgeIcon,omitempty"`
		ShowAdditionalCluster     bool  `json:"showAdditionalCluster,omitempty"`
		RingStatus                any   `json:"ringStatus,omitempty"`
		PrimarySubtitle           struct {
			TextDirection                 string   `json:"textDirection"`
			Text                          string   `json:"text"`
			AttributesV2                  []any    `json:"attributesV2"`
			AccessibilityTextAttributesV2 []any    `json:"accessibilityTextAttributesV2"`
			AccessibilityText             any      `json:"accessibilityText"`
			RecipeTypes                   []string `json:"$recipeTypes"`
			Type                          string   `json:"$type"`
		} `json:"primarySubtitle,omitempty"`
		BadgeText                any    `json:"badgeText,omitempty"`
		TrackingID               string `json:"trackingId,omitempty"`
		ActorNavigationURL       any    `json:"actorNavigationUrl,omitempty"`
		AddEntityToSearchHistory bool   `json:"addEntityToSearchHistory,omitempty"`
		Summary                  any    `json:"summary,omitempty"`
		Image                    struct {
			Attributes []struct {
				ScalingType any `json:"scalingType"`
				DetailData  struct {
					ProfilePictureWithoutFrame     any `json:"profilePictureWithoutFrame"`
					ProfilePictureWithRingStatus   any `json:"profilePictureWithRingStatus"`
					CompanyLogo                    any `json:"companyLogo"`
					Icon                           any `json:"icon"`
					SystemImage                    any `json:"systemImage"`
					NonEntityGroupLogo             any `json:"nonEntityGroupLogo"`
					VectorImage                    any `json:"vectorImage"`
					NonEntityProfessionalEventLogo any `json:"nonEntityProfessionalEventLogo"`
					ProfilePicture                 any `json:"profilePicture"`
					ImageURL                       any `json:"imageUrl"`
					ProfessionalEventLogo          any `json:"professionalEventLogo"`
					NonEntityCompanyLogo           any `json:"nonEntityCompanyLogo"`
					NonEntitySchoolLogo            any `json:"nonEntitySchoolLogo"`
					GroupLogo                      any `json:"groupLogo"`
					SchoolLogo                     any `json:"schoolLogo"`
					GhostImage                     any `json:"ghostImage"`
					NonEntityProfilePicture        struct {
						Profile     string   `json:"*profile"`
						RingStatus  any      `json:"ringStatus"`
						RecipeTypes []string `json:"$recipeTypes"`
						VectorImage any      `json:"vectorImage"`
						Type        string   `json:"$type"`
					} `json:"nonEntityProfilePicture"`
				} `json:"detailData"`
				TintColor          any      `json:"tintColor"`
				RecipeTypes        []string `json:"$recipeTypes"`
				TapTargets         []any    `json:"tapTargets"`
				DisplayAspectRatio any      `json:"displayAspectRatio"`
				Type               string   `json:"$type"`
			} `json:"attributes"`
			ActionTarget                any      `json:"actionTarget"`
			AccessibilityTextAttributes []any    `json:"accessibilityTextAttributes"`
			TotalCount                  any      `json:"totalCount"`
			AccessibilityText           any      `json:"accessibilityText"`
			RecipeTypes                 []string `json:"$recipeTypes"`
			Type                        string   `json:"$type"`
		} `json:"image,omitempty"`
		LazyLoadedActions any `json:"lazyLoadedActions,omitempty"`
		SecondarySubtitle struct {
			TextDirection                 string   `json:"textDirection"`
			Text                          string   `json:"text"`
			AttributesV2                  []any    `json:"attributesV2"`
			AccessibilityTextAttributesV2 []any    `json:"accessibilityTextAttributesV2"`
			AccessibilityText             any      `json:"accessibilityText"`
			RecipeTypes                   []string `json:"$recipeTypes"`
			Type                          string   `json:"$type"`
		} `json:"secondarySubtitle,omitempty"`
		NavigationURL          string `json:"navigationUrl,omitempty"`
		EntityEmbeddedObject   any    `json:"entityEmbeddedObject,omitempty"`
		UnreadIndicatorDetails any    `json:"unreadIndicatorDetails,omitempty"`
		Target                 any    `json:"target,omitempty"`
		ActorTrackingUrn       any    `json:"actorTrackingUrn,omitempty"`
		NavigationContext      struct {
			OpenExternally bool     `json:"openExternally"`
			RecipeTypes    []string `json:"$recipeTypes"`
			URL            string   `json:"url"`
			Type           string   `json:"$type"`
		} `json:"navigationContext,omitempty"`
		LazyLoadedActions0 string `json:"*lazyLoadedActions,omitempty"`
	} `json:"included"`
}

type githubResponseOfSearchingForUsers struct {
	TotalCount        int  `json:"total_count"`
	IncompleteResults bool `json:"incomplete_results"`
	Items             []struct {
		Login             string  `json:"login"`
		ID                int     `json:"id"`
		NodeID            string  `json:"node_id"`
		AvatarURL         string  `json:"avatar_url"`
		GravatarID        string  `json:"gravatar_id"`
		URL               string  `json:"url"`
		HTMLURL           string  `json:"html_url"`
		FollowersURL      string  `json:"followers_url"`
		FollowingURL      string  `json:"following_url"`
		GistsURL          string  `json:"gists_url"`
		StarredURL        string  `json:"starred_url"`
		SubscriptionsURL  string  `json:"subscriptions_url"`
		OrganizationsURL  string  `json:"organizations_url"`
		ReposURL          string  `json:"repos_url"`
		EventsURL         string  `json:"events_url"`
		ReceivedEventsURL string  `json:"received_events_url"`
		Type              string  `json:"type"`
		SiteAdmin         bool    `json:"site_admin"`
		Score             float64 `json:"score"`
	} `json:"items"`
}
type Employee struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

func getResponseFromFile(filename string, optionalArg ...string) []byte {
	rawReq, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	proxyURL, err := url.Parse("http://127.0.0.1:8080")
	if err != nil {
		fmt.Println("Error parsing proxy URL:", err)
		// return
	}

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		Proxy:           http.ProxyURL(proxyURL),
	}

	// Create a new client with the transport
	client := &http.Client{
		Transport: transport,
	}
	regex := regexp.MustCompile("start:[^,]+")
	// Parse the raw HTTP request
	reqParts := strings.SplitN(string(rawReq), "\r\n\r\n", 2)
	reqLine := strings.Split(strings.TrimSpace(reqParts[0]), " ")
	method := reqLine[0]
	url := reqLine[1]
	url = "https://www.linkedin.com" + url
	if len(optionalArg) > 0 {
		// url = strings.Replace(url, "username", optionalArg[0], 1)
		url = regex.ReplaceAllString(url, "start:"+optionalArg[0])
	} else {
		url = regex.ReplaceAllString(url, "start:0")
	}
	body := ""
	if len(reqParts) > 1 {
		body = reqParts[1]
	}

	// Create a new HTTP request
	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		panic(err)
	}

	// Parse the headers from the raw HTTP request and add them to the new request
	headers := strings.Split(reqParts[0], "\n")[1:]
	for _, header := range headers {
		headerParts := strings.SplitN(header, ":", 2)
		req.Header.Add(strings.TrimSpace(headerParts[0]), strings.TrimSpace(headerParts[1]))
	}

	// Send the HTTP request and print the response
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		// return
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return respBody
}
func sendHTTPRequest(start string) []Employee {
	var employees []Employee

	var reqBody2 response
	if err := json.Unmarshal(getResponseFromFile("testdir/request.txt", start), &reqBody2); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	for _, includedRec := range reqBody2.Included {
		if includedRec.Title.Text != "null" && includedRec.Title.Text != "" && !strings.Contains(includedRec.Title.Text, "Member") && !strings.Contains(includedRec.Title.Text, "Founder") && !strings.Contains(includedRec.Title.Text, "CEO") && !strings.Contains(includedRec.Title.Text, "CTO") && !strings.Contains(includedRec.Title.Text, "COO") && !strings.Contains(includedRec.Title.Text, "CFO") && !strings.Contains(includedRec.Title.Text, "CIO") && !strings.Contains(includedRec.Title.Text, "CPO") && !strings.Contains(includedRec.Title.Text, "CMO") && !strings.Contains(includedRec.Title.Text, "CDO") && !strings.Contains(includedRec.Title.Text, "CRO") && !strings.Contains(includedRec.Title.Text, "CSO") && !strings.Contains(includedRec.Title.Text, "CLO") && !strings.Contains(includedRec.Title.Text, "might benefit") {

			employeeName := includedRec.Title.Text
			employeeLocation := includedRec.SecondarySubtitle.Text

			employees = append(employees, Employee{Name: employeeName, Location: employeeLocation})
		}

	}
	// return employeeJSON
	return employees
}

type Employees struct {
	Employees []Employee `json:"employees"`
}
type githubUserInfo struct {
	Login             string    `json:"login"`
	ID                int       `json:"id"`
	NodeID            string    `json:"node_id"`
	AvatarURL         string    `json:"avatar_url"`
	GravatarID        string    `json:"gravatar_id"`
	URL               string    `json:"url"`
	HTMLURL           string    `json:"html_url"`
	FollowersURL      string    `json:"followers_url"`
	FollowingURL      string    `json:"following_url"`
	GistsURL          string    `json:"gists_url"`
	StarredURL        string    `json:"starred_url"`
	SubscriptionsURL  string    `json:"subscriptions_url"`
	OrganizationsURL  string    `json:"organizations_url"`
	ReposURL          string    `json:"repos_url"`
	EventsURL         string    `json:"events_url"`
	ReceivedEventsURL string    `json:"received_events_url"`
	Type              string    `json:"type"`
	SiteAdmin         bool      `json:"site_admin"`
	Name              string    `json:"name"`
	Company           any       `json:"company"`
	Blog              string    `json:"blog"`
	Location          string    `json:"location"`
	Email             any       `json:"email"`
	Hireable          any       `json:"hireable"`
	Bio               string    `json:"bio"`
	TwitterUsername   string    `json:"twitter_username"`
	PublicRepos       int       `json:"public_repos"`
	PublicGists       int       `json:"public_gists"`
	Followers         int       `json:"followers"`
	Following         int       `json:"following"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func getGithubUser(token string, username string) (githubUserInfo, error) {
	url := "https://api.github.com/users/" + username

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return githubUserInfo{}, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return githubUserInfo{}, err
	}
	defer resp.Body.Close()

	var responseBody githubUserInfo
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		return githubUserInfo{}, err
	}
	time.Sleep(500 * time.Millisecond)
	return responseBody, nil
}

func searchUsers(token string, username string) (githubResponseOfSearchingForUsers, error) {
	url := "https://api.github.com/search/users?q=" + username

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return githubResponseOfSearchingForUsers{}, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return githubResponseOfSearchingForUsers{}, err
	}
	defer resp.Body.Close()

	var responseBody githubResponseOfSearchingForUsers
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		return githubResponseOfSearchingForUsers{}, err
	}
	time.Sleep(500 * time.Millisecond)
	return responseBody, nil

}
func generateLocationVariations(location string) []string {
	var rev_slc []string
	// split the location string into separate words
	words := strings.Split(location, " ")

	// create a slice to store the variations
	var variations []string

	// loop through each level of the pyramid
	for i := 1; i <= len(words); i++ {
		// create a slice to hold the words for this level of the pyramid
		levelWords := make([]string, i)

		// copy the first i words into the slice
		copy(levelWords, words[:i])

		// trim any commas from the end of the words
		// for j := range levelWords {

		// 	levelWords[j] = strings.TrimSuffix(levelWords[j], ",")

		// }

		// add the level to the variations slice
		wholeOne := strings.Join(levelWords, " ")
		wholeOne = strings.TrimSuffix(wholeOne, ",")
		variations = append(variations, wholeOne)
		// variations = reverseStringSlice(variations)
		// rev_slc := []int{}
		for i := range variations {
			// reverse the order
			rev_slc = append(rev_slc, variations[len(variations)-1-i])
		}
	}

	return rev_slc
}
func isInSlice(s string, slice []string) bool {
	for _, item := range slice {
		if strings.Contains(s, item) {
			return true
		}
	}
	return false
}

func getUserReposDetails(username string, token string) string {

	url := "https://api.github.com/users/" + username + "/repos"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("error")
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error")
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error")
	}

	time.Sleep(500 * time.Millisecond)
	return string(responseBody)

}

func getURLResponse(url string, authToken string) (string, error) {
	// Create new HTTP GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	// Add authorization header to request
	req.Header.Add("Authorization", "Bearer "+authToken)

	// Send HTTP request with client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Convert response body to string
	response := string(body)
	time.Sleep(500 * time.Millisecond)
	return response, nil
}

type githubsearchresults struct {
	TotalCount        int  `json:"total_count"`
	IncompleteResults bool `json:"incomplete_results"`
	Items             []struct {
		Name       string `json:"name"`
		Path       string `json:"path"`
		Sha        string `json:"sha"`
		URL        string `json:"url"`
		GitURL     string `json:"git_url"`
		HTMLURL    string `json:"html_url"`
		Repository struct {
			ID       int    `json:"id"`
			NodeID   string `json:"node_id"`
			Name     string `json:"name"`
			FullName string `json:"full_name"`
			Private  bool   `json:"private"`
			Owner    struct {
				Login             string `json:"login"`
				ID                int    `json:"id"`
				NodeID            string `json:"node_id"`
				AvatarURL         string `json:"avatar_url"`
				GravatarID        string `json:"gravatar_id"`
				URL               string `json:"url"`
				HTMLURL           string `json:"html_url"`
				FollowersURL      string `json:"followers_url"`
				FollowingURL      string `json:"following_url"`
				GistsURL          string `json:"gists_url"`
				StarredURL        string `json:"starred_url"`
				SubscriptionsURL  string `json:"subscriptions_url"`
				OrganizationsURL  string `json:"organizations_url"`
				ReposURL          string `json:"repos_url"`
				EventsURL         string `json:"events_url"`
				ReceivedEventsURL string `json:"received_events_url"`
				Type              string `json:"type"`
				SiteAdmin         bool   `json:"site_admin"`
			} `json:"owner"`
			HTMLURL          string `json:"html_url"`
			Description      any    `json:"description"`
			Fork             bool   `json:"fork"`
			URL              string `json:"url"`
			ForksURL         string `json:"forks_url"`
			KeysURL          string `json:"keys_url"`
			CollaboratorsURL string `json:"collaborators_url"`
			TeamsURL         string `json:"teams_url"`
			HooksURL         string `json:"hooks_url"`
			IssueEventsURL   string `json:"issue_events_url"`
			EventsURL        string `json:"events_url"`
			AssigneesURL     string `json:"assignees_url"`
			BranchesURL      string `json:"branches_url"`
			TagsURL          string `json:"tags_url"`
			BlobsURL         string `json:"blobs_url"`
			GitTagsURL       string `json:"git_tags_url"`
			GitRefsURL       string `json:"git_refs_url"`
			TreesURL         string `json:"trees_url"`
			StatusesURL      string `json:"statuses_url"`
			LanguagesURL     string `json:"languages_url"`
			StargazersURL    string `json:"stargazers_url"`
			ContributorsURL  string `json:"contributors_url"`
			SubscribersURL   string `json:"subscribers_url"`
			SubscriptionURL  string `json:"subscription_url"`
			CommitsURL       string `json:"commits_url"`
			GitCommitsURL    string `json:"git_commits_url"`
			CommentsURL      string `json:"comments_url"`
			IssueCommentURL  string `json:"issue_comment_url"`
			ContentsURL      string `json:"contents_url"`
			CompareURL       string `json:"compare_url"`
			MergesURL        string `json:"merges_url"`
			ArchiveURL       string `json:"archive_url"`
			DownloadsURL     string `json:"downloads_url"`
			IssuesURL        string `json:"issues_url"`
			PullsURL         string `json:"pulls_url"`
			MilestonesURL    string `json:"milestones_url"`
			NotificationsURL string `json:"notifications_url"`
			LabelsURL        string `json:"labels_url"`
			ReleasesURL      string `json:"releases_url"`
			DeploymentsURL   string `json:"deployments_url"`
		} `json:"repository"`
		Score float64 `json:"score"`
	} `json:"items"`
}

func appendToFile(filename string, text string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(text)
	return err
}

func threadingInLinkedInEmployees(employee Employee, eligbleUsers []string, locationMethod bool, userKeywords []string, keywordsMethod bool, outputToFile bool, githubToken string, optionalArgs ...string) {
	var outputLocation string

	if len(optionalArgs) > 0 && outputToFile {
		outputLocation = optionalArgs[0]
	}
	foundKeyword := false
	encodedUserName := url.QueryEscape(employee.Name)
	if employee.Name != "" || len(employee.Name) > 0 {
		color.Cyan("[+] Searching For: " + employee.Name)
	}
	foundUsers, err := searchUsers(githubToken, encodedUserName)
	locationsGeneratedFromLinkedIn := generateLocationVariations(employee.Location)
	if err != nil {
		color.Red("[-] Can not search users")
	}
	for i := 0; i < len(foundUsers.Items); i++ {
		// fmt.Println("  [+] Testing user: ", foundUsers.Items[i].Login)
		user := foundUsers.Items[i]
		userInformaiton, err := getGithubUser(githubToken, user.Login)

		if err != nil {
			color.Red("[-] Can not get user information")
		}
		if locationMethod {
			userLocation := userInformaiton.Location

			if isInSlice(userLocation, locationsGeneratedFromLinkedIn) {
				color.Green("[*] Found: " + user.Login)
				if outputToFile {
					appendToFile(outputLocation, user.Login+"\n")
				}
			}
		} else if keywordsMethod {
			userReposInfo := getUserReposDetails(user.Login, githubToken)

			for _, keyword := range userKeywords {
				if foundKeyword {
					break
				} else if strings.Contains(keyword, userReposInfo) {
					color.Green("[*] Found: " + user.Login + ", keyword: " + keyword)
					// eligbleUsers = append(eligbleUsers, user.Login)
					if outputToFile {
						appendToFile(outputLocation, user.Login+"\n")
					}
					foundKeyword = true
					break
				} else {
					stringOfUserInfo, err := json.Marshal(userInformaiton)
					if err != nil {
						fmt.Println(stringOfUserInfo)
						color.Red("[-] Can not get user repos information #0")
					}

					if strings.Contains(string(stringOfUserInfo), keyword) {
						color.Green("[*] Found: " + user.Login + ", keyword: " + keyword)
						if outputToFile {
							appendToFile(outputLocation, user.Login+"\n")
						}
						foundKeyword = true
						break
					}
					searchResults, err := getURLResponse("https://api.github.com/search/code?q=user:"+user.Login+"+"+keyword, githubToken)
					if err != nil {
						color.Red("[-] https://api.github.com/search/code?q=user:" + user.Login + "+" + keyword)
						color.Red("[-] Can not get user repos information #1")
					}

					var githubsearchresultsobject githubsearchresults
					err = json.Unmarshal([]byte(searchResults), &githubsearchresultsobject)
					if err != nil {
						{
						}
					}
					if githubsearchresultsobject.TotalCount > 0 {
						color.Green("[*] Found: " + user.Login + ", keyword: " + keyword)
						foundKeyword = true
						if outputToFile {
							appendToFile(outputLocation, user.Login+"\n")
						}
						break
					}

				}
			}
		}

	}
}
func main() {

	color.Green("\n\t\t                                    /$$$$$$           ")
	color.Green("\t\t                                   /$$$_  $$          ")
	color.Green("\t\t /$$$$$$/$$$$  /$$   /$$ /$$   /$$| $$$$\\ $$ /$$   /$$")
	color.Green("\t\t| $$_  $$_  $$| $$  | $$|  $$ /$$/| $$ $$ $$|  $$ /$$/")
	color.Green("\t\t| $$ \\ $$ \\ $$| $$  | $$ \\  $$$$/ | $$\\ $$$$ \\  $$$$/ ")
	color.Green("\t\t| $$ | $$ | $$| $$  | $$  >$$  $$ | $$ \\ $$$  >$$  $$ ")
	color.Green("\t\t| $$ | $$ | $$|  $$$$$$/ /$$/\\  $$|  $$$$$$/ /$$/\\  $$")
	color.Green("\t\t|__/ |__/ |__/ \\______/ |__/  \\__/ \\______/ |__/  \\__/")
	fmt.Println()
	color.Green("\t\t[+] mulef - LinkedIn Employee Finder - v1.0")
	color.Green("\t\t[+] github@mux0x")
	fmt.Println()
	keywords := flag.String("keywords", "", "comma-separated list of keywords")
	mode := flag.String("mode", "", "mode of finding employees (location, keywords)")
	requestFile := flag.String("LinkedInRequest", "", "path of the linkedin request file")
	githubToken := flag.String("token", "", "github token")
	outputLocation := flag.String("output", "", "path of the output file")
	var outputToFile bool
	flag.Parse()

	if flag.Lookup("keywords") == nil {
		color.Red("[-] Keywords flag not specified")
		flag.Usage()
		os.Exit(1)
	}
	if flag.Lookup("mode") == nil {
		color.Red("[-] Mode flag not specified")
		flag.Usage()
		os.Exit(1)
	}
	if flag.Lookup("LinkedInRequest") == nil {
		color.Red("[-] LinkedInRequest flag not specified")
		flag.Usage()
		os.Exit(1)
	}
	if flag.Lookup("token") == nil {
		color.Red("[-] Token flag not specified")
		flag.Usage()
		os.Exit(1)
	}

	if flag.Lookup("output") != nil {
		outputToFile = true
	} else {
		outputToFile = false
	}

	allEmployees := Employees{}
	var locationMethod bool
	var keywordsMethod bool

	if *mode == "location" {
		locationMethod = true
	} else if *mode == "keywords" {
		keywordsMethod = true
	} else {
		color.Red("[-] Invalid mode")
		os.Exit(1)
	}

	var reqBody2 response
	if err := json.Unmarshal(getResponseFromFile(*requestFile), &reqBody2); err != nil { // Parse []byte to the go struct pointer
		color.Red("[-] Can not unmarshal JSON")
	}

	employeesCount := reqBody2.Data.Data.SearchDashClustersByAll.Metadata.TotalResultCount

	numPages := int(math.Ceil(float64(employeesCount) / 10))
	// fmt.Println("rounds: ", numPages)
	color.Cyan("[+] Processing LinkedIn Request")
	for i := 0; i <= numPages; i++ {
		startIndex := i * 10
		// fmt.Println("start: ", startIndex)
		newEmployee := sendHTTPRequest(strconv.Itoa(startIndex))
		allEmployees.Employees = append(allEmployees.Employees, newEmployee...)

	}
	userKeywords := strings.Split(*keywords, ",")
	// fmt.Println("all employees: ", allEmployees.Employees)

	// maxThreads := 1
	threadCount := 1

	done := make(chan bool)
	employeeChan := make(chan Employee)

	for i := 0; i < threadCount; i++ {
		go func() {
			for {
				select {
				case e := <-employeeChan:
					threadingInLinkedInEmployees(e, userKeywords, locationMethod, userKeywords, keywordsMethod, outputToFile, *githubToken, *outputLocation)
				case <-done:
					return
				}
			}
		}()
	}

	for _, e := range allEmployees.Employees {
		employeeChan <- e
	}

	close(employeeChan)

	for i := 0; i < threadCount; i++ {
		done <- true
	}

}
