package usecase

import (
	"log/slog"

	"github.com/sisoputnfrba/tp-golang/cpu/internal/core/gateway/rest/dto"
)

func CheckInterrupt(interrupt *dto.Interrupt) {
	if interrupt.TID == dto.CtxEnEjecucion.TID && interrupt.PID == dto.CtxEnEjecucion.PID {
		//slog.Info("CHECK-INTERRUPT", "PID EXEC:", dto.CtxEnEjecucion.PID, "TID EXEC", dto.CtxEnEjecucion.TID)
		//slog.Info("CHECK-INTERRUPT", "PID INT:", interrupt.PID, "TID INT", interrupt.TID)
		dto.Interrumpir = true
		dto.PidInterrupt = interrupt.PID
		dto.TidInterrupt = interrupt.TID
	} else {
		slog.Info("CHECK_INTERRUPT: El TID a interrumpir no coincide con el TID en ejecución y no se detiene la ejecucion")
	}
}
