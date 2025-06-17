package grpc

import (
	"github.com/google/uuid"
	db "github.com/Lykeion-org/lexora-dictionary-service/internal/db"
	pb "github.com/Lykeion-org/lexora-dictionary-service/internal/grpc/generated"
)

// Proto to DB conversions
func convertProtoToReferent(r *pb.Referent) *db.Referent {
	if r == nil {
		return nil
	}
	
	uid, _ := uuid.Parse(r.GetUid())
	return &db.Referent{
		UID:         uid,
		EnReference: r.GetEnReference(),
		ImageSource: stringPtr(r.GetImageSource()),
	}
}

func convertProtoToSymbol(s *pb.Symbol) *db.Symbol {
	if s == nil {
		return nil
	}

	uid, _ := uuid.Parse(s.GetUid())
	var lemmaUID *uuid.UUID
	if s.GetLemma() != nil {
		uid, _ := uuid.Parse(s.GetLemma().GetUid())
		lemmaUID = &uid
	}

	return &db.Symbol{
		UID:        uid,
		Language:   int(s.GetLanguage()),
		SymbolType: int(s.GetSymbolType()),
		LemmaUID:   lemmaUID,
	}
}

func convertProtoToWord(w *pb.Word) *db.Word {
	if w == nil {
		return nil
	}

	uid, _ := uuid.Parse(w.GetUid())
	return &db.Word{
		UID:         uid,
		Word:        w.GetWord(),
		SoundSource: stringPtr(w.GetSoundSource()),
		IPA:         stringPtr(w.GetIpa()),
		WordType:    int(w.GetWordType()),
	}
}

// DB to Proto conversions
func convertReferentToProto(r *db.Referent) *pb.Referent {
	if r == nil {
		return nil
	}

	var symbols []*pb.Symbol
	for _, s := range r.Symbols {
		symbols = append(symbols, convertSymbolToProto(&s))
	}

	return &pb.Referent{
		Uid:         r.UID.String(),
		EnReference: r.EnReference,
		ImageSource: derefString(r.ImageSource),
		Symbols:     symbols,
	}
}

func convertSymbolToProto(s *db.Symbol) *pb.Symbol {
	if s == nil {
		return nil
	}

	return &pb.Symbol{
		Uid:        s.UID.String(),
		Language:   pb.Language(s.Language),
		SymbolType: pb.SymbolType(s.SymbolType),
		Lemma:      convertWordToProto(s.Lemma),
		Words:      convertWordListToProto(s.Words),
	}
}

func convertWordToProto(w *db.Word) *pb.Word {
	if w == nil {
		return nil
	}

	return &pb.Word{
		Uid:         w.UID.String(),
		Word:        w.Word,
		SoundSource: derefString(w.SoundSource),
		Ipa:         derefString(w.IPA),
		WordType:    pb.WordType(w.WordType),
	}
}

func convertWordListToProto(wl []db.Word) []*pb.Word {
	var words []*pb.Word
	for _, w := range wl {
		words = append(words, convertWordToProto(&w))
	}
	return words
}

// Helper functions
func derefString(ptr *string) string {
	if ptr != nil {
		return *ptr
	}
	return ""
}

func stringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
