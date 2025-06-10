package grpc

import (
	"context"
	"errors"
	db "github.com/Lykeion/lexora-dictionary-service/internal/db"
	pb "github.com/Lykeion/lexora-dictionary-service/internal/grpc/generated"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type LanguageServer struct {
	pb.UnimplementedLanguageServiceServer

	ReferentSvc *db.ReferentService
	SymbolSvc *db.SymbolService
	WordSvc *db.WordService
}

// READ
func (s *LanguageServer) GetReferent(ctx context.Context, req *pb.GetReferentRequest) (*pb.Referent, error){
	//convert string to uid
	uid, err := uuid.Parse(req.GetUid()); if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid UID: %v", err)
	}

	ref, err := s.ReferentSvc.GetReferent(ctx, uid); if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "referent not found")
		}
		return nil, status.Errorf(codes.Internal, "db error: %v", err)
	}

	return convertReferentToProto(ref), nil
}

func (s *LanguageServer) GetSymbol(ctx context.Context, req *pb.GetSymbolRequest) (*pb.Symbol, error){
	uid, err := uuid.Parse(req.GetUid()); if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid UID: %v", err)
	}

	sym, err := s.SymbolSvc.GetSymbol(ctx, uid); if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "symbol not found")
		}
		return nil, status.Errorf(codes.Internal, "db error: %v", err)
	}

	return convertSymbolToProto(sym), nil

}
func (s *LanguageServer) GetWord(ctx context.Context, req *pb.GetWordRequest) (*pb.Word, error){
	uid, err := uuid.Parse(req.GetUid()); if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid UID: %v", err)
	}

	wrd, err := s.WordSvc.GetWord(ctx, uid); if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "word not found")
		}
		return nil, status.Errorf(codes.Internal, "db error: %v", err)
	}

	return convertWordToProto(wrd), nil
}
func (s *LanguageServer) ListReferents(ctx context.Context, req *pb.ListReferentsRequest) (*pb.ListReferentsResponse, error)
func (s *LanguageServer) FindReferents(ctx context.Context, req *pb.FindReferentsRequest) (*pb.FindReferentsResponse, error)

// CREATE
func (s *LanguageServer) CreateReferent(ctx context.Context, req *pb.CreateReferentRequest) (*pb.Referent, error){
	var ref *db.Referent = &db.Referent{
		UID: uuid.New(),
		EnReference: req.EnReference,
		ImageSource: req.ImageSource,
	}
	
	createdRef, err := s.ReferentSvc.CreateReferent(ctx, ref); if err != nil {
		return nil, status.Errorf(codes.Internal, "db: error: %v", err)
	}

	return convertReferentToProto(createdRef), nil
}

func (s *LanguageServer) CreateSymbol(ctx context.Context, req *pb.CreateSymbolRequest) (*pb.Symbol, error) {
	
	refUid, err := uuid.Parse(req.GetReferentUid()); if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid referent UID: %v", err)
	}

	var sym *db.Symbol = &db.Symbol{
		UID: uuid.New(),
		ReferentUID: refUid,
		Language: int(req.GetLanguage()),
		SymbolType: int(req.GetSymbolType()),
	}
	
	createdSym, err := s.SymbolSvc.CreateSymbol(ctx, sym); if err != nil {
		return nil, status.Errorf(codes.Internal, "db: error: %v", err)
	}

	return convertSymbolToProto(createdSym), nil
}

func (s *LanguageServer) CreateWord(ctx context.Context, req *pb.CreateWordRequest) (*pb.Word, error) {

	var wrd *db.Word = &db.Word{
		UID: uuid.New(),
		Word: req.GetWord(),
		WordType: int(req.GetWordType()),
	}
	
	createdWrd, err := s.WordSvc.CreateWord(ctx, wrd); if err != nil {
		return nil, status.Errorf(codes.Internal, "db: error: %v", err)
	}

	return convertWordToProto(createdWrd), nil

}

// UPDATE
func (s *LanguageServer) UpdateReferent(ctx context.Context, req *pb.UpdateReferentRequest) (*pb.Referent, error) {
	if req.GetReferent() == nil {
		return nil, status.Error(codes.InvalidArgument, "referent is required")
	}

	// First get existing referent
	uid, err := uuid.Parse(req.GetReferent().GetUid())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid UID: %v", err)
	}

	existing, err := s.ReferentSvc.GetReferent(ctx, uid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "referent not found")
		}
		return nil, status.Errorf(codes.Internal, "db error: %v", err)
	}

	// Update fields
	existing.EnReference = req.GetReferent().GetEnReference()
	if req.GetReferent().GetImageSource() != "" {
		imgSrc := req.GetReferent().GetImageSource()
		existing.ImageSource = &imgSrc
	} else {
		existing.ImageSource = nil
	}

	// Save changes
	if err := s.ReferentSvc.UpdateReferent(ctx, existing); err != nil {
		return nil, status.Errorf(codes.Internal, "db error: %v", err)
	}

	return convertReferentToProto(existing), nil
}
func (s *LanguageServer) UpdateSymbol(ctx context.Context, req *pb.UpdateSymbolRequest) (*pb.Symbol, error) {
	if req.GetSymbol() == nil {
		return nil, status.Error(codes.InvalidArgument, "symbol is required")
	}

	// First get existing symbol
	uid, err := uuid.Parse(req.GetSymbol().GetUid())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid UID: %v", err)
	}

	existing, err := s.SymbolSvc.GetSymbol(ctx, uid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "symbol not found")
		}
		return nil, status.Errorf(codes.Internal, "db error: %v", err)
	}

	// Update fields
	existing.Language = int(req.GetSymbol().GetLanguage())
	existing.SymbolType = int(req.GetSymbol().GetSymbolType())
	// Note: ReferentUID cannot be updated directly through Symbol proto

	// Save changes
	if err := s.SymbolSvc.UpdateSymbol(ctx, existing); err != nil {
		return nil, status.Errorf(codes.Internal, "db error: %v", err)
	}

	return convertSymbolToProto(existing), nil
}
func (s *LanguageServer) UpdateWord(ctx context.Context, req *pb.UpdateWordRequest) (*pb.Word, error) {
	if req.GetWord() == nil {
		return nil, status.Error(codes.InvalidArgument, "word is required")
	}

	// First get existing word
	uid, err := uuid.Parse(req.GetWord().GetUid())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid UID: %v", err)
	}

	existing, err := s.WordSvc.GetWord(ctx, uid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "word not found")
		}
		return nil, status.Errorf(codes.Internal, "db error: %v", err)
	}

	// Update fields
	existing.Word = req.GetWord().GetWord()
	existing.WordType = int(req.GetWord().GetWordType())
	if req.GetWord().GetSoundSource() != "" {
		soundSrc := req.GetWord().GetSoundSource()
		existing.SoundSource = &soundSrc
	} else {
		existing.SoundSource = nil
	}
	if req.GetWord().GetIpa() != "" {
		ipa := req.GetWord().GetIpa()
		existing.IPA = &ipa
	} else {
		existing.IPA = nil
	}

	// Save changes
	if err := s.WordSvc.UpdateWord(ctx, existing); err != nil {
		return nil, status.Errorf(codes.Internal, "db error: %v", err)
	}

	return convertWordToProto(existing), nil
}

// DELETE
func (s *LanguageServer) DeleteReferent(ctx context.Context, req *pb.DeleteReferentRequest) (*pb.Referent, error) {
	uid, err := uuid.Parse(req.GetReferentUid())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid UID: %v", err)
	}

	// First get the referent to return it after deletion
	ref, err := s.ReferentSvc.GetReferent(ctx, uid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "referent not found")
		}
		return nil, status.Errorf(codes.Internal, "db error: %v", err)
	}

	// Delete the referent
	if err := s.ReferentSvc.DeleteReferent(ctx, uid); err != nil {
		return nil, status.Errorf(codes.Internal, "db error: %v", err)
	}

	return convertReferentToProto(ref), nil
}
func (s *LanguageServer) DeleteSymbol(ctx context.Context, req *pb.DeleteSymbolRequest) (*pb.Symbol, error) {
	uid, err := uuid.Parse(req.GetSymbolUid())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid UID: %v", err)
	}

	// First get the symbol to return it after deletion
	sym, err := s.SymbolSvc.GetSymbol(ctx, uid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "symbol not found")
		}
		return nil, status.Errorf(codes.Internal, "db error: %v", err)
	}

	// Delete the symbol
	if err := s.SymbolSvc.DeleteSymbol(ctx, uid); err != nil {
		return nil, status.Errorf(codes.Internal, "db error: %v", err)
	}

	return convertSymbolToProto(sym), nil
}
func (s *LanguageServer) DeleteWord(ctx context.Context, req *pb.DeleteWordRequest) (*pb.Word, error) {
	uid, err := uuid.Parse(req.GetWordUid())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid UID: %v", err)
	}

	// First get the word to return it after deletion
	wrd, err := s.WordSvc.GetWord(ctx, uid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "word not found")
		}
		return nil, status.Errorf(codes.Internal, "db error: %v", err)
	}

	// Delete the word
	if err := s.WordSvc.DeleteWord(ctx, uid); err != nil {
		return nil, status.Errorf(codes.Internal, "db error: %v", err)
	}

	return convertWordToProto(wrd), nil
}

// LINKING OPERATIONS
func (s *LanguageServer) LinkSymbolToReferent(context.Context, *pb.LinkSymbolToReferentRequest) (*pb.LinkSymbolToReferentResponse, error)
func (s *LanguageServer) LinkWordToSymbol(context.Context, *pb.LinkWordToSymbolRequest) (*pb.LinkWordToSymbolResponse, error)
func (s *LanguageServer) SetSymbolLemma(context.Context, *pb.SetSymbolLemmaRequest) (*pb.SetSymbolLemmaResponse, error)
