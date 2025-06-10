package grpc

import (
	db "github.com/Lykeion/lexora-dictionary-service/internal/db"
	pb "github.com/Lykeion/lexora-dictionary-service/internal/grpc/generated"
)


func convertWordToProto(w *db.Word) *pb.Word {
	return &pb.Word{
		Uid: w.UID.String(),
		Word: w.Word,
		SoundSource: derefString(w.SoundSource),
		Ipa: derefString(w.IPA),
		WordType: pb.WordType(w.WordType),
	}
}

func convertWordListToProto(wl []db.Word) []*pb.Word {
	
}

func convertSymbolToProto(s *db.Symbol) *pb.Symbol {
	return &pb.Symbol{
		Uid:  s.UID.String(),
		Language: pb.Language(s.Language),
		SymbolType: pb.SymbolType(s.SymbolType),
		Lemma: convertWordToProto(s.Lemma),
		Words: convertWordListToProto(s.Words),
	}
}

func convertReferentToProto(r *db.Referent) *pb.Referent {

}

func derefString(ptr *string) string {
	if ptr != nil {
		return *ptr
	}
	return ""
}