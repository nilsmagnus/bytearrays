import java.nio.ByteBuffer

public static byte[] bytes(long l) {
    ByteBuffer.allocate(Long.BYTES).putLong(l).array();
}

def pretty(byte[] b) {
    def sb = new StringBuilder()
    for(i=0;i<b.length;i++){
        sb.append(String.format("\\x%02X", b[i]))
    }
    sb.toString()
}

def long2hex(l) {pretty(bytes(l))}

def now = System.currentTimeMillis()
println now
println long2hex(now)
