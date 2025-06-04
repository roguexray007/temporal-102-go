package durableexecution

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"go.temporal.io/sdk/activity"
	// TODO Add the import here, needed to use the Activity logger
)

func TranslateTerm(ctx context.Context, input TranslationActivityInput) (TranslationActivityOutput, error) {
	// TODO Define an Activity logger
	logger := activity.GetLogger(ctx)

	// TODO log Activity invocation, at the Info level, and include the term being
	//      translated and the language code as name-value pairs
	logger.Info("TranslateTerm", "term", input.Term, "languageCode", input.LanguageCode)

	lang := url.QueryEscape(input.LanguageCode)
	term := url.QueryEscape(input.Term)
	url := fmt.Sprintf("http://localhost:9998/translate?lang=%s&term=%s", lang, term)

	resp, err := http.Get(url)
	if err != nil {
		return TranslationActivityOutput{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return TranslationActivityOutput{}, err
	}

	// This string will contain either the translated term, if the service could
	// perform the translation, or the error message, if it was unsuccessful
	content := string(body)

	status := resp.StatusCode
	if status >= 400 {
		// This means that we successfully called the service, but it could not
		// perform the translation for some reason
		return TranslationActivityOutput{},
			fmt.Errorf("HTTP Error %d: %s", status, content)
	}

	// TODO  use the Debug level to log the successful translation and include the
	//       translated term as a name-value pair
	logger.Debug("TranslateTerm", "translation", content)

	output := TranslationActivityOutput{
		Translation: content,
	}

	return output, nil
}
