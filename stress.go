package stressapi

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"time"

	stress "github.com/ikawaha/kagome-stress/gen/stress"
	"github.com/ikawaha/kagome/tokenizer"
)

// stress service example implementation.
// The example methods log the requests and return zero values.
type stresssrvc struct {
	logger    *log.Logger
	tokenizer tokenizer.Tokenizer
	cancel    func()
	mux       sync.Mutex
}

// NewStress returns the stress service implementation.
func NewStress(logger *log.Logger) stress.Service {
	return &stresssrvc{
		logger:    logger,
		tokenizer: tokenizer.NewWithDic(tokenizer.SysDicUni()),
		cancel:    nil,
	}
}

func (s *stresssrvc) stress(ctx context.Context) {
	log.Print("start!!!")
	ticker := time.Tick(10 * time.Second)
	for {
		select {
		case <-ctx.Done():
			log.Print("stop!!!")
			return
		case <-ticker:
			log.Print("*")
		default:
			go func() {
				if tokens := s.tokenizer.Tokenize("すもももももももものうち"); len(tokens) != 9 {
					log.Printf("%+v", tokens)
				}
			}()
			go func() {
				if tokens := s.tokenizer.Tokenize("シルバニア"); len(tokens) != 4 && tokens[0].Surface != "シルバ" {
					log.Printf("%+v", tokens)
				}
			}()
			go func() {
				txt := "キラキラヒカルサイフヲダシテキ ラキラヒカルサカナヲカツタキラ キラヒカルオンナモカツタキラキ ラヒカルサカナヲカツテキラキラ ヒカルオナベニイレテキラキラヒ カルオンナガモツテキラキラヒカ ルオナベノサカナキラキラヒカル オツリノオカネキラキラヒカルオ ンナトフタリキラキラヒカルサカ ナヲモツテキラキラヒカルオカネ ヲモツテキラキラヒカルヨミチヲ カエルキラキラヒカルホシゾラダ ツタキラキラヒカルナミダヲダシ テキラキラヒカルオンナガナイタ"
				if tokens := s.tokenizer.Tokenize(txt); len(tokens) != 29 {
					log.Printf("len=%d, %+v", len(tokens), tokens)
				}
			}()
			time.Sleep(time.Duration(rand.Intn(500)) * time.Microsecond)
		}
	}
}

// Start implements start.
func (s *stresssrvc) Start(ctx context.Context) (err error) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.logger.Print("stress.start")
	if s.cancel != nil {
		s.logger.Print("already run")
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	go s.stress(ctx)
	s.cancel = cancel
	return nil
}

// Stop implements stop.
func (s *stresssrvc) Stop(ctx context.Context) (err error) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.logger.Print("stress.stop")
	if s.cancel == nil {
		s.logger.Print("not running yet")
		return
	}
	s.cancel()
	s.cancel = nil
	return
}

// Tokenize implements tokenize.
func (s *stresssrvc) Tokenize(ctx context.Context, p *stress.TokenizePayload) (res stress.TokenCollection, err error) {
	s.logger.Print("stress.tokenize")
	tokens := s.tokenizer.Tokenize(p.Sentence)
	for _, v := range tokens {
		res = append(res, &stress.Token{
			Surface: v.Surface,
			Pos:     v.Pos(),
			Start:   v.Start,
			End:     v.End,
			Type:    v.Class.String(),
		})
	}
	return res, nil
}
